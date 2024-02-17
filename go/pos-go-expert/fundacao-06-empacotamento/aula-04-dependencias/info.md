
# help

## iniciar um modulo

$ `go mod init github.com/taranttini/study/go/pos-go-expert/fundacao-06-empacotamento/aula-04-dependencias`

## importante

é importante usar um path único, e de preferencia uma url que exista esse código (no caso um repositório)

# detalhes

quando criado uma variavel, extruct, ... inciando com letra maiúscula, ela se torna acessivel, caso eu inicie com letra minúscula o item não fica acessível


## depois de iniciar o modulo

é importante executar o `go mod tidy` para iniciar e baixar os pacotes para o projeto


## obs

caso o projeto tenha dependencias nao usada, ou removida, o go tidy irá remover esses itens, o go cria um arquivo go.sum, onde irá registrar informações a respeito dos pacotes utilizados

## outra forma de obter um pacote

é possível usar o comando `go get PATH_DO_PACOTE` a ser baixado, nesse caso como queremos o google uuid iremos usar o comando  `go get github.com/google/uuid` e ai ele irá alimentar o go.mod e go.sum