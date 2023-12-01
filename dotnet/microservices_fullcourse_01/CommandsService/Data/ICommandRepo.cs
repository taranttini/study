using CommandsService.Models;

namespace CommandsService;

public interface ICommandRepo
{
    bool SaveChanges();

    // Platforms
    IEnumerable<Platform> GetAllPlatforms();
    void CreatePlatform(Platform platform);
    bool PlaftormExists(int platformId);
    bool ExternalPlatormExist(int externalPlataformId);

    // Commands
    IEnumerable<Command> GetCommandsForPlatform(int platformId);
    Command? GetCommand(int platformId, int commandId);
    void CreateCommand(int platformId, Command command);

}
