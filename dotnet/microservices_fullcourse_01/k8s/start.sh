minikube start

#minikube addons enable ingress
#minikube addons enable ingress-dns

kubectl create secret generic mssql --from-literal=SA_PASSWORD="my@KEY#123_yourPass" 

kubectl apply -f mssql-plat-depl.yaml 

kubectl apply -f local-pvc.yaml

kubectl apply -f rabbitmq-depl.yaml

kubectl apply -f platforms-depl.yaml

kubectl apply -f commands-depl.yaml

##kubectl apply -f platforms-np-srv.yaml # nao necessario

minikube addons enable ingress
#minikube addons enable ingress-dns

kubectl apply -f ingress-srv.yaml


minikube tunnel

# tentando rodar local
# kubectl proxy --address='127.0.0.1' --disable-filter=true


# export DOTNET_NUGET_SIGNATURE_VERIFICATION=false

# minikube dashboard