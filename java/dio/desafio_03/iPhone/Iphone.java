package study.java.dio.desafio_03.iPhone;

import study.java.dio.desafio_03.iPhone.Aplicativo.AparelhoTelefonico;
import study.java.dio.desafio_03.iPhone.Aplicativo.NavegadorInternet;
import study.java.dio.desafio_03.iPhone.Aplicativo.ReprodutorMusical;

public class Iphone {
    private ReprodutorMusical reprodutorMusical;
    private NavegadorInternet navegadorInternet;
    private AparelhoTelefonico aparelhoTelefonico;
    
    public Iphone() {
    }
    
    public void voltaParaHome() {
        System.out.println("voltar para a home");
    }
    
    public void rotacionarTela() {
        System.out.println("rotacionar tela");
    }
    
    public ReprodutorMusical abrirReprodutorMusical() {
        if (this.reprodutorMusical == null) {
            this.reprodutorMusical = new ReprodutorMusical();
        }
        return this.reprodutorMusical;
    }
    
    public NavegadorInternet abrirNavegadorInternet() {
        if (this.navegadorInternet == null) {
            this.navegadorInternet = new NavegadorInternet();
        }
        return this.navegadorInternet;
    }
    
    public AparelhoTelefonico abrirAparelhoTelefonico() {
        if (this.aparelhoTelefonico == null) {
            this.aparelhoTelefonico = new AparelhoTelefonico();
        }
        return this.aparelhoTelefonico;
    }
}
