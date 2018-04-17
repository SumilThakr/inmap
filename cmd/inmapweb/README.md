# InMAP web

This is a web interface for the InMAP model and related tools.

Docker build steps:
* Build the image: `docker build -t inmapweb .`
* Run a container: `docker run --publish 8080:10000 --name inmapwebtest --rm inmapweb`
* The webpage is at: `https://localhost:8080`
* Stop the container: `docker stop inmapwebtest`

Deploying to Kubernetes:
* https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app
