dotnet dev-certs https --trust

dotnet dev-certs https -v



#### podman

podman login docker.io

podman build -t taranttini/platformservice .

### excecutar imagem
podman run -p 8080:80 -d taranttini/platformservice

### inicial container
docker start CONTAINER_ID

### parar container
docker stop CONTAINER_ID

### subir imagem para o docker
podman push taranttini/platformservice



### kind

kind get clusters

kind create cluster --config mykind.yaml

/c/dev/kubectl cluster-info --context kind-mykind

kind delete cluster -n mykind



/c/dev/kubectl cluster-info dump

/c/dev/kubectl api-resources

####

/c/dev/kubectl apply -f platforms-depl.yaml

/c/dev/kubectl apply -f platforms-np-srv.yaml

/c/dev/kubectl get all

/c/dev/kubectl get nodes
/c/dev/kubectl get pods
/c/dev/kubectl get service

# /c/dev/kubectl port-forward service/platformservice-service 80
/c/dev/kubectl port-forward service/platformservice-srv 80

/c/dev/kubectl delete deployment platforms-depl