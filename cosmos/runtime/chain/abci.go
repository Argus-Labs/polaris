// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package chain

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/rs/zerolog/log"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/beacon/engine"

	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/core/types"
)

func (wbc *WrappedBlockchain) ProcessProposal(
	ctx sdk.Context, req *abci.RequestProcessProposal,
) (*abci.ResponseProcessProposal, error) {
	var (
		err error
	)

	// We have to run the PreBlocker && BeginBlocker to get the chain into the state
	// it'll be in when the EVM transaction actually runs.
	if _, err = wbc.app.PreBlocker(ctx, &abci.RequestFinalizeBlock{
		Txs:                req.Txs,
		Time:               req.Time,
		Misbehavior:        req.Misbehavior,
		Height:             req.Height,
		NextValidatorsHash: req.NextValidatorsHash,
		ProposerAddress:    req.ProposerAddress,
	}); err != nil {
		return &abci.ResponseProcessProposal{
			Status: abci.ResponseProcessProposal_REJECT,
		}, err
	} else if _, err = wbc.app.BeginBlocker(ctx); err != nil {
		return &abci.ResponseProcessProposal{
			Status: abci.ResponseProcessProposal_REJECT,
		}, err
	}

	ctx.GasMeter().RefundGas(ctx.GasMeter().GasConsumed(), "process proposal")
	ctx.BlockGasMeter().RefundGas(ctx.BlockGasMeter().GasConsumed(), "process proposal")
	ctx = ctx.WithKVGasConfig(storetypes.GasConfig{}).
		WithTransientKVGasConfig(storetypes.GasConfig{}).
		WithGasMeter(storetypes.NewInfiniteGasMeter())

	// Pull an execution payload out of the proposal.
	var envelope *engine.ExecutionPayloadEnvelope
	for _, tx := range req.Txs {
		var sdkTx sdk.Tx
		sdkTx, err = wbc.app.TxDecode(tx)
		if err != nil {
			// should have been verified in prepare proposal, we
			// ignore it for now (i.e VE extensions will fail decode).
			continue
		}

		log.Info().Msgf("looping over %d sdk.Msg's", len(sdkTx.GetMsgs()))
		for _, sdkMsg := range sdkTx.GetMsgs() {
			if env, ok := sdkMsg.(*evmtypes.WrappedPayloadEnvelope); ok {
				envelope = env.UnwrapPayload()
				break
			}
		}
	}

	// If the proposal doesn't contain an ethereum envelope, we should just move on.
	if envelope == nil {
		return &abci.ResponseProcessProposal{
			Status: abci.ResponseProcessProposal_ACCEPT,
		}, nil
	} else {
		log.Info().Msg("success!! FOUND ENVELOPE")
	}

	// Convert it to a block.
	var block *types.Block
	if block, err = engine.ExecutableDataToBlock(*envelope.ExecutionPayload, nil, nil); err != nil {
		ctx.Logger().Error("failed to build evm block", "err", err)
		return &abci.ResponseProcessProposal{
			Status: abci.ResponseProcessProposal_REJECT,
		}, err
	}

	// Insert the block into the chain.
	receipts, err := wbc.InsertBlockWithoutSetHead(ctx, block)
	if err != nil {
		ctx.Logger().Error("failed to insert block", "err", err)
		return &abci.ResponseProcessProposal{
			Status: abci.ResponseProcessProposal_REJECT,
		}, err
	}

	if wbc.hook != nil {
		txs := block.Transactions()
		log.Info().Msgf("Running post-block hooks for %d txs", len(txs))
		header := block.Header()
		signer := types.MakeSigner(wbc.Blockchain.Config(), header.Number, header.Time)
		wbc.hook(block.Transactions(), receipts, signer)
	}

	return &abci.ResponseProcessProposal{
		Status: abci.ResponseProcessProposal_ACCEPT,
	}, nil
}
