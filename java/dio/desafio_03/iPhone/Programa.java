package study.java.dio.desafio_03.iPhone;

public class Programa {
    public static void main(String[] args) {

        Iphone iphone = new Iphone();

        var reprodutorMusical = iphone.abrirReprodutorMusical();
        reprodutorMusical.tocar();
        reprodutorMusical.pausa();
        reprodutorMusical.telaScroll();
        reprodutorMusical.selecionaMusica("parabéns pra você");
        reprodutorMusical.tocar();
        reprodutorMusical.pausa();
        iphone.rotacionarTela();
        iphone.voltaParaHome();

        var navegadorInternet = iphone.abrirNavegadorInternet();
        navegadorInternet.exibirPagina("http://google.com");
        navegadorInternet.atualizarAba();
        navegadorInternet.telaScroll();
        navegadorInternet.adicionarNovaAba();
        navegadorInternet.exibirPagina("http://bing.com");
        navegadorInternet.fecharAba();
        iphone.rotacionarTela();
        iphone.voltaParaHome();

        var aparelhoTelefonico = iphone.abrirAparelhoTelefonico();
        aparelhoTelefonico.ligar("11 1234 1234");
        aparelhoTelefonico.encerrarLigacao();
        aparelhoTelefonico.atender();
        aparelhoTelefonico.encerrarLigacao();
        aparelhoTelefonico.telaRotacionar();
        aparelhoTelefonico.iniciarCorreioDeVoz();
        aparelhoTelefonico.encerrarCorreioDeVoz();
        iphone.voltaParaHome();

    }
}
