package br.com.dio.desafio.dominio;

public class Curso extends Conteudo{

    private int cargaHoraria;

    @Override
    public double calcularXp() {
        return XP_PADRAO * cargaHoraria;
    }

    public Curso() {
    }


    public int getCargaHoraria() {
        return cargaHoraria;
    }

    public void setCargaHoraria(int cargaHoraria) {
        this.cargaHoraria = cargaHoraria;
    }

    @Override
    public String toString() {
        return "\n\tCurso: {" +
                "\n\t\t \"titulo\": \"" + getTitulo() + "\"" +
                ",\n\t\t \"descricao\": \"" + getDescricao() + "\"" +
                ",\n\t\t \"cargaHoraria\": " + cargaHoraria +
                "\n\t}";
    }
}
