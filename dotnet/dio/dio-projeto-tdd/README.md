# Projeto TDD

## Criando a Solução e arquivos iniciais do projeto

Criando a solução

`dotnet new sln -o projeto-tdd`

Acessando a pasta projeto-tdd, e criando o projeto de lib do serviço de agenda

`dotnet new classlib -o ServicoAgenda`

Incluindo na solution o projeto de lib do serviço de agenda

`dotnet sln add ./ServicoAgenda/ServicoAgenda.csproj`

Criando o projeto de teste do serviço de agenda

`dotnet new xunit -o ServicoAgenda.Tests`

Incluindo na solution o projeto de teste do serviço de agenda

`dotnet sln add ./ServicoAgenda.Tests/ServicoAgenda.Tests.csproj`

Incluíndo no projeto de teste do serviço de agenda, a referência do projeto de lib do serviço de agenda

`dotnet add ./ServicoAgenda.Tests/ServicoAgenda.Tests.csproj reference ./ServicoAgenda/ServicoAgenda.csproj

## o que foi feito

criação dos teste e criação das regras

rodando os teste, de dentro da pasta projeto-tdd

`dotnet test -l "console;verbosity=normal"`
