using System.Diagnostics;
using MethodTimer;
using Microsoft.AspNetCore.Mvc;

namespace performing_api.Controllers;

[ApiController]
[Route("[controller]")]
public class WeatherForecastController : ControllerBase
{
    private static readonly string[] Summaries = new[]
    {
        "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
    };

    private readonly ILogger<WeatherForecastController> _logger;

    public WeatherForecastController(ILogger<WeatherForecastController> logger)
    {
        _logger = logger;
    }

    [HttpGet]
    [Route("/GetLoggerStopwatch")]
    public IEnumerable<WeatherForecast> GetLoggerStopwatch()
    {
        var sw = Stopwatch.StartNew();
        try
        {
            return Enumerable.Range(1, 5).Select(index => new WeatherForecast
                {
                    Date = DateOnly.FromDateTime(DateTime.Now.AddDays(index)),
                    TemperatureC = Random.Shared.Next(-20, 55),
                    Summary = Summaries[Random.Shared.Next(Summaries.Length)]
                })
                .ToArray();
        }
        finally
        {
            sw.Stop();
            Console.WriteLine($"Took {sw.Elapsed}ms");
        }
    }

    [Time("Retrieved weather for {days}")]
    [HttpGet]
    [Route("/GetLoggerTimer")]
    public async Task<IEnumerable<WeatherForecast>> GetLoggerTimer([FromQuery] int days = 5)
    {
        await Task.Delay(Random.Shared.Next(0, days));
        return Enumerable.Range(1, days).Select(index => new WeatherForecast
            {
                Date = DateOnly.FromDateTime(DateTime.Now.AddDays(index)),
                TemperatureC = Random.Shared.Next(-20, 55),
                Summary = Summaries[Random.Shared.Next(Summaries.Length)]
            })
            .ToArray();
    }
}