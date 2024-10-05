package study.java.dio.desafio_03.iPhone.Aplicativo;

public class AparelhoTelefonico extends Aplicativo implements Funcionalidade {
    public AparelhoTelefonico() {
        super.setNome("Aparelho Telefonico");
        logarAcao("inicia");
    }

    boolean ligacaoAtiva;
    boolean correioDeVozAtivo;

    @Override
    public void telaRotacionar() {
        logarAcao("tela rotacionar - funcionalidade sem acao");
    }

    @Override
    public void telaScroll() {
        logarAcao("scrollar tela");
    }

    public void iniciarCorreioDeVoz() {
        logarAcao("fechar aba");
    }

    public void atender() {
        logarAcao("atender");
        ligacaoAtiva = true;
    }

    public void ligar(String numero) {
        logarAcao("ligar '" + numero + "'");
    }

    public void encerrarLigacao() {
        logarAcao("encerrar ligação");
        if (ligacaoAtiva == false) {
            System.out.println(" >> Não existe ligação ativa");
        }
        ligacaoAtiva = false;
    }

    public void encerrarCorreioDeVoz() {
        logarAcao("encerrar ligação");
        if (correioDeVozAtivo == false) {
            System.out.println(" >> Não existe correio de voz ativo");
        }
        correioDeVozAtivo = false;
    }
}
