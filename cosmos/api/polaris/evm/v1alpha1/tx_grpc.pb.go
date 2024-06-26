
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: polaris/evm/v1alpha1/tx.proto

package evmv1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MsgService_EthTransaction_FullMethodName         = "/polaris.evm.v1alpha1.MsgService/EthTransaction"
	MsgService_ProcessPayloadEnvelope_FullMethodName = "/polaris.evm.v1alpha1.MsgService/ProcessPayloadEnvelope"
)

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgServiceClient interface {
	// EthTransaction defines a method submitting Ethereum transactions.
	EthTransaction(ctx context.Context, in *WrappedEthereumTransaction, opts ...grpc.CallOption) (*WrappedEthereumTransactionResult, error)
	// ProcessPayloadEnvelope defines a method to process CL paylods.
	ProcessPayloadEnvelope(ctx context.Context, in *WrappedPayloadEnvelope, opts ...grpc.CallOption) (*WrappedPayloadEnvelopeResponse, error)
}

type msgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgServiceClient(cc grpc.ClientConnInterface) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) EthTransaction(ctx context.Context, in *WrappedEthereumTransaction, opts ...grpc.CallOption) (*WrappedEthereumTransactionResult, error) {
	out := new(WrappedEthereumTransactionResult)
	err := c.cc.Invoke(ctx, MsgService_EthTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ProcessPayloadEnvelope(ctx context.Context, in *WrappedPayloadEnvelope, opts ...grpc.CallOption) (*WrappedPayloadEnvelopeResponse, error) {
	out := new(WrappedPayloadEnvelopeResponse)
	err := c.cc.Invoke(ctx, MsgService_ProcessPayloadEnvelope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
// All implementations must embed UnimplementedMsgServiceServer
// for forward compatibility
type MsgServiceServer interface {
	// EthTransaction defines a method submitting Ethereum transactions.
	EthTransaction(context.Context, *WrappedEthereumTransaction) (*WrappedEthereumTransactionResult, error)
	// ProcessPayloadEnvelope defines a method to process CL paylods.
	ProcessPayloadEnvelope(context.Context, *WrappedPayloadEnvelope) (*WrappedPayloadEnvelopeResponse, error)
	mustEmbedUnimplementedMsgServiceServer()
}

// UnimplementedMsgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (UnimplementedMsgServiceServer) EthTransaction(context.Context, *WrappedEthereumTransaction) (*WrappedEthereumTransactionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthTransaction not implemented")
}
func (UnimplementedMsgServiceServer) ProcessPayloadEnvelope(context.Context, *WrappedPayloadEnvelope) (*WrappedPayloadEnvelopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessPayloadEnvelope not implemented")
}
func (UnimplementedMsgServiceServer) mustEmbedUnimplementedMsgServiceServer() {}

// UnsafeMsgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServiceServer will
// result in compilation errors.
type UnsafeMsgServiceServer interface {
	mustEmbedUnimplementedMsgServiceServer()
}

func RegisterMsgServiceServer(s grpc.ServiceRegistrar, srv MsgServiceServer) {
	s.RegisterService(&MsgService_ServiceDesc, srv)
}

func _MsgService_EthTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WrappedEthereumTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).EthTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgService_EthTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).EthTransaction(ctx, req.(*WrappedEthereumTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ProcessPayloadEnvelope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WrappedPayloadEnvelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ProcessPayloadEnvelope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgService_ProcessPayloadEnvelope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ProcessPayloadEnvelope(ctx, req.(*WrappedPayloadEnvelope))
	}
	return interceptor(ctx, in, info, handler)
}

// MsgService_ServiceDesc is the grpc.ServiceDesc for MsgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MsgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "polaris.evm.v1alpha1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EthTransaction",
			Handler:    _MsgService_EthTransaction_Handler,
		},
		{
			MethodName: "ProcessPayloadEnvelope",
			Handler:    _MsgService_ProcessPayloadEnvelope_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "polaris/evm/v1alpha1/tx.proto",
}
