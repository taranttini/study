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