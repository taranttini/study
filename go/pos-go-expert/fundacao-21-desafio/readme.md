
printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

go mod tidy

--

go run github.com/99designs/gqlgen init

go mod tidy

--

go run server.go

--

go run github.com/99designs/gqlgen generate

https://plataforma.fullcycle.com.br/courses/c2957fa4-1e88-4425-be86-5a17ad2664ca/302/190/177/conteudos?capitulo=177&conteudo=9875

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