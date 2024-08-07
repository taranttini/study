
```sh

# graphQL

# instalar
printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
# carregar modulos
go mod tidy

# iniciar
go run github.com/99designs/gqlgen init
# carregar modulos
go mod tidy

# verificar ser
go run server.go

# gerar
go run github.com/99designs/gqlgen generate

## gRPC
# Install

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

sudo apt install -y protobuf-compiler

# export GOPATH=$HOME/go
# export PATH=$PATH:$GOPATH/bin


# para criar dentro da pasta pb os arquivos e entidades
protoc --go_out=. --go-grpc_out=. proto/order_item.proto



// protoc --go_out=. internal/infra/pb  --go-grpc_out=. internal/infra/proto/order_item.proto
// protoc --go_out=. --go-grpc_out=. ./internal/infra/proto/order_item.proto

// evans --path ./proto --proto order_item.proto --port 50051
// evans repl --proto internal/infra/proto/order_item.proto --port 50051



# exemplo docker

FROM node:12-alpine
RUN apk add --no-cache protoc

https://plataforma.fullcycle.com.br/courses/c2957fa4-1e88-4425-be86-5a17ad2664ca/302/190/177/conteudos?capitulo=177&conteudo=9875

https://plataforma.fullcycle.com.br/courses/c2957fa4-1e88-4425-be86-5a17ad2664ca/302/190/177/conteudos?capitulo=177&conteudo=9928
```

# Clean Architect

Olá devs!

Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.

Esta listagem precisa ser feita com:

- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL

Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.

Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.