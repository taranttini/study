package study.java.dio.desafio_03.iPhone.Aplicativo;

public class NavegadorInternet extends Aplicativo implements Funcionalidade {
    public NavegadorInternet() {
        super.setNome("Navegador Internet");
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

    public void atualizarAba() {
        logarAcao("atualizar aba");        
    }

    public void fecharAba() {
        logarAcao("fechar aba");
    }

    public void adicionarNovaAba() {
        logarAcao("adicionar nova aba");
    }

    public void exibirPagina(String url) {
        logarAcao("exibir página url '" + url + "'");
    }
}
