# GO-Backend-Deployment
Design a basic model for Backend with GIN Golang API  - Build an image by a Dockerfile for this Backend, pull this image to Docker Hub and write a YAML configuration to deploy Backend into K8s by minikube.
``
Ensure the following:

1. Go is installed with version >= 1.20
2. This project is build using VS Code, you can use any editor of your choice.

``  
To deploy into K8s:
1. Pull an image: (first install docker engine)
```bash
cd demo
docker pull tutu1008/my-backend:latest
docker images
```
3. Deploy K8s (first install minikube)
```bash
minikube start
kubectl apply -f k8s-deployment.yml
kubectl get all -n default
minikube service go-service
```
5. To get API response
Example:
```bash
192.168.10.1:7004/api/GetData
192.168.10.1:7004/api/GetQueryStringData?name=Tu&age=21
```
