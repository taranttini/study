minikube start

kubectl apply -f platforms-depl.yaml

kubectl apply -f commands-depl.yaml

kubectl apply -f platforms-np-srv.yaml

kubectl apply -f ingress-srv.yaml

kubectl create secret generic mssql --from-literal=SA_PASSWORD="pa55w0rd" 

kubectl apply -f mssql-plat-depl.yaml 

kubectl apply -f local-pvc.yaml

kubectl apply -f rabbitmq-depl.yaml

minikube tunnel

# tentando rodar local
# kubectl proxy --address='127.0.0.1' --disable-filter=true


# export DOTNET_NUGET_SIGNATURE_VERIFICATION=false

# minikube dashboard