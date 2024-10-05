# Solução

O código está localizado na pasta iPhone, e execução do programa está na classe `Programa.java`

Info: Segue a solução da atividade em formato diagrama Mermaid

***OBS**: Segue imagem, caso não seja visível no modelo Mermaid*

![](iphone_one.svg)

```mermaid
classDiagram

direction LR

class Iphone {
    +voltarParaHome()
    +rotacionarTela()
    +abrirReprodutorMusical()
    +abrirNavegadorInternet()
    +abrirAparelhoTelefonico()
}

class ReprodutorMusical {
    -musica: string
    -selecionouUmaMusica() : boolean
    +tocar() 
    +pausa() 
    +selecionaMusica(string nome)
}

class AparelhoTelefonico {
    -ligacaoAtiva : boolean
    -correioDeVozAtivo : boolean
    +iniciarCorreioDeVoz()
    +atender()
    +ligar(string numero)
    +encerrarLigacao()
    +encerrarCorreioDeVoz()
}

class NavegadorInternet {
  +exibirPagina(string url)
  +adicionarNovaAba()
  +atualizarAba()
  +fecharAba()
}

class Aplicativo {
    <<abstract>>
    -nome: string
    +getNome() : string
    +setNome(String nome)
    +logarAcao(String acao)
}

class Funcionalidade {
    <<interface>>
    +telaRotacionar()
    +telaScroll()
}

Iphone ..> ReprodutorMusical : usa
Iphone ..> AparelhoTelefonico : usa
Iphone ..> NavegadorInternet : usa

%%implements
ReprodutorMusical --|> Funcionalidade 
AparelhoTelefonico --|> Funcionalidade 
NavegadorInternet --|> Funcionalidade 

%%extends
ReprodutorMusical --|> Aplicativo 
AparelhoTelefonico --|> Aplicativo 
NavegadorInternet --|> Aplicativo 

```