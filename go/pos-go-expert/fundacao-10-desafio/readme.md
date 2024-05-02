# Info

Esse documento serve para guiar como executar o projeto e validar o desafio

Desafio contido no arquivo [desafio](desafio.md)

# Execução

No diretório raiz do projeto, onde se encontra o arquivo `go.mod`, executar o comando no terminal

$ `go mod tidy`

No diretório `configs`, existe o arquivo `.env` com as configuracões padrão do projeto.

*Caso não exista, necessário criar ele com esses dados:*

```ini
HTTP_TIMEOUT = 1000 # in milisseconds
URL_BRASILAPI = https://brasilapi.com.br/api/cep/v1
URL_VIACEP = http://viacep.com.br/ws
```



Acesse o diretorio principal do código `main`

$  `cd cmd/server`

Executar o comando no terminal

$ `go run main.go`

*ou*

$ `go run main.go 02765070`

Será retornado qual **Api** trouxe a busca do cep de forma mais rápida, ou timeout por causa da requisição não concluída, ou outros erros.

exemplo de saída:

```
$ go run main.go 02765070
Consultar CEP - ViaCEP X BrasilApi

[ A api da ViaCep foi mais rápida ]
Rua José Frederico, São Paulo - SP CEP 02765-070
```