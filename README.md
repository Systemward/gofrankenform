Kubernetes deployment

# deploy the thing
kubectl create -f deployment-manifest.yaml

# figure out the service address
kubectl describe service https-server



# TODO

* Figure out how to get the rest endpoints to be accepted by the HTTPS server
* Figure out how to do auth (idc which one. Just don't spin up 502,000 terraform resources and make me go broke)
* Create the pipeline and figure out the flow
    * Deploy the https server
    * figure out the rest endpoint deployment
    * create a load balancer maybe