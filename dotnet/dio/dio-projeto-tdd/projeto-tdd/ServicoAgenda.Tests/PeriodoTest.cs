namespace ServicoAgenda.Tests;

public class PeriodoTest
{
    //[Fact]
    [Theory]
    [InlineData (20, 05, 2024)] // seg
    [InlineData (21, 05, 2024)] // ter 
    [InlineData (22, 05, 2024)] // qua
    [InlineData (23, 05, 2024)] // qui
    [InlineData (24, 05, 2024)] // sex
    public void Periodo_EhDiaDeSemana(int dia, int mes, int ano)
    {
        var periodo = new ServicoAgenda.Periodo();
        var diaDeSemana = new DateTime(ano, mes, dia); // seg, ter, qua, qui e sex

        var resultado = periodo.EhDiaDeSemana(diaDeSemana);

        Assert.True(resultado);
    }

    //[Fact]
    [Theory]
    [InlineData (25, 05, 2024)] // sab
    [InlineData (26, 05, 2024)] // dom
    public void Periodo_NaoEhDiaDeSemana(int dia, int mes, int ano)
    {
        var periodo = new ServicoAgenda.Periodo();
        var diaFinalDeSemana = new DateTime(ano ,mes , dia); // sab e dom

        var resultado = periodo.EhDiaDeSemana(diaFinalDeSemana);

        Assert.False(resultado);
    }

    //[Fact]
    [Theory]
    [InlineData (8, 00)]
    [InlineData (8, 01)]
    [InlineData (9, 30)]
    [InlineData (11, 15)]
    [InlineData (13, 30)]
    [InlineData (15, 45)]
    [InlineData (18, 59)]
    [InlineData (19, 00)]
    public void Periodo_EhHorarioComercial(int hora, int minuto)
    {
        var periodo = new ServicoAgenda.Periodo();
        var horarioComercial = new DateTime(2024, 05, 29, hora, minuto, 00); // horario entre às 8:00 até às 19:00

        var resultado = periodo.EhHorarioComercial(horarioComercial);

        Assert.True(resultado);
    }

    //[Fact]
    [Theory]
    [InlineData (5, 00)] // antes da hora
    [InlineData (7, 59)] // ultimo minuto invalido
    [InlineData (19, 01)] // primeiro minuto invalido
    [InlineData (00, 00)] // depois da hora
    public void Periodo_NaoEhHorarioComercial(int hora, int minuto)
    {
        var periodo = new ServicoAgenda.Periodo();
        var horarioNaoComercial = new DateTime(2024, 05, 29, hora, minuto, 00); // horario fora das 8:00 até às 19:00

        var resultado = periodo.EhHorarioComercial(horarioNaoComercial);

        Assert.False(resultado);
    }

    //[Fact]
    [Theory]
    [InlineData (13, 05, 2024, 08, 00)] // seg na primeira hora
    [InlineData (13, 05, 2024, 19, 00)] // seg na ultima hora
    [InlineData (17, 05, 2024, 08, 00)] // sex na primeira hora
    [InlineData (17, 05, 2024, 19, 00)] // sex na ultima hora
    public void Periodo_EhValido(int dia, int mes, int ano, int hora, int minuto)
    {
        var periodo = new ServicoAgenda.Periodo();
        var periodoValido = new DateTime(ano, mes, dia, hora, minuto, 00); // seg à sex das 8:00 até às 19:00

        var resultado = periodo.EhValido(periodoValido);

        Assert.True(resultado);
    }

    //[Fact]
    [Theory]
    [InlineData (13, 05, 2024, 07, 59)] // seg antes da hora
    [InlineData (13, 05, 2024, 19, 01)] // seg depois da hora
    [InlineData (18, 05, 2024, 12, 00)] // sab dentro da hora
    [InlineData (19, 05, 2024, 13, 00)] // dom dentro da hora
    public void Periodo_NaoEhValido(int dia, int mes, int ano, int hora, int minuto)
    {
        var periodo = new ServicoAgenda.Periodo();
        var periodoNaoEhValido = new DateTime(ano, mes, dia, hora, minuto, 00); // fora de seg à sex das 8:00 até às 19:00

        var resultado = periodo.EhValido(periodoNaoEhValido);

        Assert.False(resultado);
    }
}
