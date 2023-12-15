using CommandsService.Models;
using CommandsService.SyncDataServices.Grpc;
using Grpc.Net.Client.Configuration;

namespace CommandsService.Data;

public static class PrepDb
{
    public static void PrepPopulation(IApplicationBuilder applicationBuilder)
    {
        using var serviceScope = applicationBuilder.ApplicationServices.CreateScope();

        var grpcClient = serviceScope.ServiceProvider.GetService<IPlatformDataClient>();

        var Platforms = grpcClient.ReturnAllPlatforms();

        SeedData(serviceScope.ServiceProvider.GetService<ICommandRepo>(), Platforms);
    }

    private static void SeedData(ICommandRepo commandRepo, IEnumerable<Platform> platforms)
    {
        Console.WriteLine("Seeding new platforms...");

        foreach (var platform in platforms)
        {
            if (!commandRepo.ExternalPlatormExist(platform.ExternalID))
            {
                commandRepo.CreatePlatform(platform);
            }
            commandRepo.SaveChanges();
        }
    }
}