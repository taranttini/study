# SOLUÇÃO

Acessar a pasta do projeto go/pos-go-expert/fundacao-21-desafio

executar o comando docker-compose up

Aguardar ser criado os serviços

## API

Na pasta http, temos o arquivo `api.http` com as funções

- GET http://localhost:8080/order > que lista as ordens
- POST http://localhost:8080/order-create > que criar uma nova ordem
- POST http://localhost:8080/order-add-item > que adiciona item na ordem

## GraphQL

Temos também o arquivo `api_graph.http`

é uma cola para ser usada ao acessar o endereço `http://localhost:8080/` no navegador, e assim colar o script para rodar o GraphQL

A rotina `query queryOrderAndItem` lista as ordens

## Grpc

Não consegui em momento algum rodar o comando `evans repl -r --port 50051`, nem mesmo nas aulas do curso GoLang, mas ao rodar o comando `evans repl --proto internal/infra/proto/order_item.proto --port 50051` ai sim foi possível executar as chamadas necessárias para funcionar as rotinas do GRPC

Então para acessar o evans é necessário rodar o comando `evans repl --proto internal/infra/proto/order_item.proto --port 50051` em sua máquina local

A rotina `call ListOrders` lista as ordens

CTRL+D para encerrar o Evans+Grpc

**OBS** caso não tenha o EVANS instalado é necessário seguir o exemplo:

https://github.com/ktr0731/evans?tab=readme-ov-file#installation

### comando que me auxiliaram no processo

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
// protoc --go_out=. --go-grpc_out=. ./internal/infra/proto/order_item.proto

// evans --path ./proto --proto order_item.proto --port 50051
// evans repl --proto internal/infra/proto/order_item.proto --port 50051


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