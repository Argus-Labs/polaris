Deploying on Kubernetes


1. Localling using `minikube`

```
    # Build the account container using the minikube docker runtime.
    make build-k8s-account

    # Boot up the pods & nlb
    kubectl apply -f k8s

    # Expose the load balancer to the network
    minikube service ts-account --url

    # Check to see if it is gucci gang.
    curl http://127.0.0.1:<port>/api/health/v1/check

```