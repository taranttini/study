using System.Text;
using System.Text.Json;
using AutoMapper;
using Microsoft.Extensions.Configuration;
using PlatformService.Dtos;

namespace PlatformService.SyncDataService.Http;

public class HttpCommandDataClient : ICommandDataClient
{
    private readonly HttpClient _httpClient;
    private readonly IConfiguration _configuration;

    public HttpCommandDataClient(
        HttpClient httpClient,
        IConfiguration configuration
    )
    {
        _httpClient = httpClient;
        _configuration = configuration;
    }
    public async Task SendPlatformToCommand(PlatformReadDto platformReadDto)
    {
        var httpContent = new StringContent(
            JsonSerializer.Serialize(platformReadDto),
            Encoding.UTF8,
            "application/json"
        );

        var response = await _httpClient.PostAsync($"{_configuration["CommandService"]}", httpContent);

        if (response.IsSuccessStatusCode)
        {
            Console.WriteLine("--> Sync POST to CommandService was OK!");
            return;
        }
        Console.WriteLine("--> Sync POST to CommandService was NOT OK!");
    }
}