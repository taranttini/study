dotnet dev-certs https --trust

dotnet dev-certs https -v



#### podman

podman login docker.io

podman build -t taranttini/platformservice .

podman build -t docker.io/taranttini/commandservice .

### podman listar imagens

podman images

### podman remover imagem

podman rmi IMAGE_ID

### excecutar imagem
podman run -p 8080:80 -d taranttini/platformservice

podman run -p 8080:80 -d docker.io/taranttini/commandservice

### inicial container
podman start CONTAINER_ID

### parar container
podman stop CONTAINER_ID

### subir imagem para o docker
podman push taranttini/platformservice

podman push docker.io/taranttini/commandservice



### kind

kind get clusters

kind create cluster --config mykind.yaml

/c/dev/kubectl cluster-info --context kind-mykind

kind delete cluster -n mykind



/c/dev/kubectl cluster-info dump

/c/dev/kubectl api-resources

####

/c/dev/kubectl apply -f platforms-depl.yaml

/c/dev/kubectl apply -f commands-depl.yaml

/c/dev/kubectl apply -f platforms-np-srv.yaml

/c/dev/kubectl rollout restart deployment platforms-depl

/c/dev/kubectl get all

/c/dev/kubectl get nodes
/c/dev/kubectl get pods
/c/dev/kubectl get service
/c/dev/kubectl get deployments
/c/dev/kubectl get namespaces
/c/dev/kubectl get storageclass


# para rodar esse serviço isolado
/c/dev/kubectl port-forward service/platforms-clusterip-srv 80

# para rodar esse serviço isolado
/c/dev/kubectl port-forward service/commands-clusterip-srv 80

# para rodar ambos serviços
/c/dev/kubectl port-forward service/platformservice-srv 80

/c/dev/kubectl delete deployment platforms-depl



# acessar o link para configurar o ingress nginx

https://kubernetes.github.io/ingress-nginx/deploy/#docker-desktop

comando

/c/dev/kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.2/deploy/static/provider/aws/deploy.yaml

/c/dev/kubectl get pods --namespace=ingress-nginx

/c/dev/kubectl get service --namespace=ingress-nginx


# abrir arquivo hosts do windows

incluir a linha

`127.0.0.1 acme.com`

para que possamos chamar esse endereço e ele responder

#

/c/dev/kubectl apply -f ingress-srv.yaml