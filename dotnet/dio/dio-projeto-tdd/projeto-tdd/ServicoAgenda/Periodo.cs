using System.Linq;

namespace ServicoAgenda;

public class Periodo
{
    public bool EhDiaDeSemana(DateTime data)
    {
         DayOfWeek[] listaDiasDeSemana = [
            DayOfWeek.Monday, // sex
            DayOfWeek.Tuesday, // ter
            DayOfWeek.Wednesday, // qua
            DayOfWeek.Thursday, // qui
            DayOfWeek.Friday // sex
        ];

        return listaDiasDeSemana.Contains(data.DayOfWeek); 
    }

    public bool EhHorarioComercial(DateTime data)
    {
        var dataMinima = new DateTime(data.Year, data.Month, data.Day, 8, 0, 0);
        var dataMaxima = new DateTime(data.Year, data.Month, data.Day, 19, 0, 0);

        return dataMinima <= data && data <= dataMaxima;
    }

    public bool EhValido(DateTime data)
    {
        return EhDiaDeSemana(data) && EhHorarioComercial(data);
    }
}
