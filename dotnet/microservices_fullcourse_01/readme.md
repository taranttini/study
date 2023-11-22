# arquivos hosts

sudo nano /etc/hosts

# parar all dockers

docker kill $(docker ps -q) 

# remover all dockers containers

docker rm $(docker ps -a -q) 

# remover all dockers containers images

docker rmi $(docker images -q)

# instalar minikube

curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube

# iniciar minikube 

minikube start

# minikube status e monitoramento

https://minikube.sigs.k8s.io/docs/start/

minikube dashboard

# habilitar certificado

dotnet dev-certs https --trust

dotnet dev-certs https -v

https://youtu.be/DgVjEo3OGBI?si=19aJDNcCfVDI5B7c&t=20939

https://www.youtube.com/watch?v=n0zkkoL8eNs

https://www.youtube.com/watch?v=gMwAhKddHYQ&list=PLzYkqgWkHPKBcDIP5gzLfASkQyTdy0t4k&index=12

https://www.youtube.com/watch?v=CqCDOosvZIk

https://www.youtube.com/watch?v=Rqz9XiSqH3E


#### install kind / kubectl

https://kind.sigs.k8s.io/docs/user/quick-start/#installation

https://kubernetes.io/pt-br/docs/tasks/tools/install-kubectl-linux/

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

kubectl cluster-info --context kind-mykind

kind delete cluster -n mykind



kubectl cluster-info dump

kubectl api-resources

####

kubectl apply -f platforms-depl.yaml

kubectl apply -f commands-depl.yaml

kubectl apply -f platforms-np-srv.yaml

kubectl rollout restart deployment platforms-depl

kubectl get all

kubectl get nodes
kubectl get pods
kubectl get service
kubectl get deployments
kubectl get namespaces
kubectl get storageclass
kubectl get pv
kubectl get pvc

kubectl delete storageclass STORAGE_CLASS_NAME

## para tentar acessar algum storage
kubectl exec -it CONTAINER_NAME -- /bin/bash


# para rodar esse serviço isolado
kubectl port-forward service/platforms-clusterip-srv 80

sudo -E kubectl port-forward service/platforms-clusterip-srv 80

# para rodar esse serviço isolado
kubectl port-forward service/commands-clusterip-srv 80

sudo -E kubectl port-forward service/commands-clusterip-srv 80

# para rodar ambos serviços
kubectl port-forward service/platformservice-srv 80

sudo -E kubectl port-forward service/platformservice-srv 80

kubectl delete deployment platforms-depl



# acessar o link para configurar o ingress nginx

https://kubernetes.github.io/ingress-nginx/deploy/#docker-desktop

comando

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.2/deploy/static/provider/aws/deploy.yaml

kubectl get pods --namespace=ingress-nginx

kubectl get service --namespace=ingress-nginx


# abrir arquivo hosts do windows

incluir a linha

`127.0.0.1 acme.com`

para que possamos chamar esse endereço e ele responder

#

kubectl apply -f ingress-srv.yaml

sudo -E kubectl apply -f ingress-srv.yaml

**caso gere problemas**

kubectl get ValidatingWebhookConfiguration

kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission

# criando volume persistênte

kubectl apply -f local-pvc.yaml

## se necessário deletar

kubectl delete pvc NOME_DO_ITEM

# gerar base sql server

kubectl create secret generic mssql --from-literal=SA_PASSWORD="pa55w0rd" 

# executar o sqlserver

kubectl apply -f mssql-plat-depl.yaml 

kubectl apply -f local-pvc.yaml

sudo -E kubectl port-forward service/mssql-clusterip-srv 1433

# atualizar projeto docker

docker build -t taranttini/platformservice .

# publicar 
docker push taranttini/platformservice


 kubectl rollout restart deployments platforms-depl