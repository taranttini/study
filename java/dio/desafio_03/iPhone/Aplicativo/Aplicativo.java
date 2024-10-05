package study.java.dio.desafio_03.iPhone.Aplicativo;

public abstract class Aplicativo {
    private String nome;
    public String getNome() {
        return nome;
    }
    public void setNome(String nome) {
        this.nome = nome;
    }
    public void logarAcao(String nomeAcao) {
        System.out.println(this.getNome() + " | " + nomeAcao);
    }
}
