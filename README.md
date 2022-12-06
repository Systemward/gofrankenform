Kubernetes deployment

# deploy the thing
kubectl create -f deployment-manifest.yaml

# figure out the service address
kubectl describe service https-server
