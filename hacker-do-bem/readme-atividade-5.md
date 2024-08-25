# README

## DevSecOps - Aula 1.5: Containers e Docker

## Atividade 1.5 : Aplicação web em um ambiente Docker

Bem-vindo à nossa atividade prática! Hoje, você vai aprender a configurar e executar uma aplicação web de reserva de salas usando containers Docker no Linux.

Envie um relatório, contendo um sumário do resultado, considerando a aplicação em execução, envie print do resultado do comando "docker ps", assim como o da aplicação, mostrando algumas reservas feitas.

Formatação e Apresentação

```s

Tamanho do Documento: Máximo de 2 páginas.

Formato: PDF.
Clareza e Concisão: Apresente as informações de forma clara e concisa, evitando detalhes desnecessários.
Estrutura: Utilize a estrutura abaixo, organizando o conteúdo em seções claras e distintas.

```
Estrutura do Relatório

```s
Título: Sumário do Resultado - Aplicação Web em um Ambiente Docker
Autor: [Seu Nome]
Data: [Data de Entrega]
```
### Objetivo

Gerar e executar um sistema web de reserva de salas usando containers Docker, no Linux.

Nesta atividade, você vai aprender a:

Instalar e configurar um container Docker.

Importar uma aplicação Golang para o ambiente Docker.

Explorar alguns comandos básicos do Docker e Docker-Compose

### Pré-requisitos

Computador com acesso à internet e permissões para instalar novos softwares

Alguma familiaridade com terminais e comandos básicos do Linux.

### Passo 1: Acesso ao ambiente do Lab

a. Faça o acesso ao ambiente do Lab com seu código de acesso, conforme instruções já recebidas.

### Passo 2: Configurando o app na VM

a. Acesse via ssh a VM da atividade, por meio do terminal no endereço IP 192.168.98.10, usando a conta "aluno", senha "rnpesr". Abra o terminal e digite o comando abaixo:

```sh
ssh aluno@192.168.98.10
```
b. Crie um diretorio chamado "reserva-salas" com o comando "mkdir /home/aluno/reserva-salas".

```sh
mkdir /home/aluno/reserva-salas
```
c. Copie todo o conteúdo do diretório /curso/reserva-salas para o diretório acima recém-criado e depois vá para o diretório reserva-salas.

```sh
cd /home/aluno
cp -r curso/aula5/reserva-salas/* ./reserva-salas/
cd ./reserva-salas
```
d. O código da aplicação reservas, escrita em golang, está no diretório /home/reserva-salas/reservas.go.

e. O código da página de login da aplicação, está no diretório /home/reserva-salas/login.html

f. O código do menu da aplicação, está no diretório /home/reserva-salas/menu.html

g. O código para fazer as reservas, está no diretório /home/reserva-salas/reserva.html

h. Aqui está o código para cancelar uma reserva, está no diretório /home/reserva-salas/cancela.html

h O código para ver o estado de uma reserva, está no diretório /home/reserva-salas/status.html

i. O diretório /home/aluno/reserva-salas/Dockerfile contém o arquivo Docker com as instruções para gerar a imagem do container, copiando todos os arquivos do código fonte, compilando, gerando o executável e depois rodando o sistema.

j. Dê o comando "cd /home/aluno/reserva-salas/dados" e veja o arquivo salas.csv, que contém o nome das salas que podem ser reservadas.

```s
Sala A
Sala B
```

j. Agora, nessa mesma pasta de dados, também existe o arquivo users.csv, que contém os usuários autorizados a usar o sistema. O arquivo contém o nome e senha do usuário, assim como se é ou não um usuário com poderes de administrador do sistema.

```s
user1,pass1,false
user2,pass2,false
admin,admin123,true
```

l. Assim, vimos todo o código fonte da aplicação de reservas, assim como o Dockerfile para gerar a imagem do container.

### Passo 3: Gerando o container

Dê o comando "cd /home/aluno/reserva-salas e gere a imagem do container com o comando abaixo:

```sh
docker build -t reserva-salas .
```
Após isso, a imagem do container foi gerada, com o nome "reserva-salas", e pode ser verificada através do comando abaixo

```sh
docker images
```

Tendo sido gerada a imagem, vamos executa-la, mapeando a porta 8080 do container para a porta 8080 do host:

```sh
docker run -d -p 8080:8080 --rm --name reserva-salas reserva-salas
```

Pode-se verificar que o container está em execução, usando o comando abaixo:

```sh
docker ps
```
Você deve ver o container sistema-reserva-salas listado na saída do comando. Salve a saida desse comando para depois ser enviado como tarefa realizada.

Para usar o sistema de reserva de salas, que está rodando em um container Docker, na sua maquina física, navegue para http://192.168.98.10:8080.

Entre no sistema de reserva, usando um dos usuários autorizados (listados no arquivo /home/aluno/reserva-salas/dados/users.csv) .

Navegue por suas paginas, testando as opções e reservando diferentes salas. Lembre que o cancelamento de reservas já feitas só pode ser realizada por um usuário com poder de administração. Tire um screenshot da tela de reservas, mostrando algumas reservas feitas, para depois ser enviado como tarefa realizada.

Para terminar o container, use o comando "docker stop reserva-salas".

### Passo 4: Uso do Docker-Compose

No passo anterior, usamos os comandos docker cli individualmente para dar build, gerar a imagem e rodar o container.

Outra possibilidade seria o uso do docker-compose para fazer tudo isso a partir de um único arquivo de configuração.

Copie o arquivo docker-compose1.yml para docker-compose.yml.

Verifique o conteúdo do arquivo /home/aluno/reserva-salas/docker-compose.yml:

```yaml
version: '3'
services:
 web:
  build:
    context: .
    dockerfile: Dockerfile
  image: reserva-salas
  container_name: reserva-salas
  ports:
    - 8080:8080
```

Execução do docker-compose:

a. No terminal, dê o comando "cd /home/aluno/reserva-salas" e execute o comando abaixo para gerar a imagem do container:

```sh
docker-compose build
```

a. Depois, dê o seguinte comando para iniciar o container, conforme definido no arquivo docker-compose.yml:

```sh
docker-compose up -d
```

Verifique com o comando "docker ps" se o container está em execução.

A partir da sua maquina física, acesse http://192.168.98.10:8080 no navegador para verificar se o sistema de reserva de salas está rodando corretamente.

Execute o comando "docker-compose down" para remover o container e desabilitar o sistema de reserva de salas.

Verifique com o comando "docker ps" que o container não está mais em execução.

### Passo 5: Persistência de dados

Usando o docker-compose, rode novamente o container e use o sistema de reservas para fazer algumas reservas de salas.

Pare o container e depois reinicie-o novamente. Ao usar novamente o sistema de salas, você vai verificar que as reservas feitas na etapa anterior foram perdidas.

Isso ocorreu porque as reservas foram guardadas no próprio container, e ao ser parado, todo seu conteúdo foi perdido

Para criar uma persistência dos dados, pode-se usar o recurso de volume do docker.

Copie o arquivo docker-compose2.yml para docker-compose.yml.

Confirme que o arquivo /home/aluno/reserva-salas/docker-compose.yml tem o conteúdo abaixo

```yaml
version: '3'
services:
 web:
  build:
    context: .
    dockerfile: Dockerfile
  image: reserva-salas
  container_name: reserva-salas
  ports:
    - 8080:8080
  volumes:
    - /home/aluno/reserva-salas/dados:/app/dados
Verifique que ao final do docker-compose.yml foi adicionado a diretiva volume, mapeando o diretório /app/dados no container para o diretório /home/aluno/reserva-salas/dados na VM.

Assim, os arquivos no diretório /app/dados estarão mapeados também no diretório do host, e não serão perdidos se o container for destruído.

Para confirmar, rode o "docker-compose up -d" para iniciar o container, use o sistema de reservas e faça algumas reservas.

Depois, pare o container e o reinicie, via docker-compose. Você verá que as reservas feitas anteriormente agora não foram perdidas, mostrando que há persistência dos dados agora.