package study.java.dio.desafio_03.iPhone.Aplicativo;

public class ReprodutorMusical extends Aplicativo implements Funcionalidade {

    public ReprodutorMusical() {
        super.setNome("Reprodutor Musical");
        logarAcao("inicia");
    }

    @Override
    public void telaRotacionar() {
        logarAcao("tela rotacionar");
    }

    @Override
    public void telaScroll() {
        logarAcao("scrollar tela");
    }

    private String musica = null;
    private boolean selecionouUmaMusica() {
        if (musica == null) {
            System.out.println(" >> selecione uma música ");
            return false;
        }
        return true;
    }

    public void tocar() {
        logarAcao("tocar");
        if (!selecionouUmaMusica()) {
            return;
        }
        System.out.println(" >> reproduzindo música: " + musica );
    }

    public void pausa() {
        logarAcao("pausa");
        if (!selecionouUmaMusica()) {
            return;
        }
        System.out.println(" >> pausou música: " + musica );
    }

    public void selecionaMusica(String nome) {
        this.musica = nome;
        logarAcao("seleciona música '" + musica + "'");
    }
}
