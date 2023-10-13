using PlatformService.Models;

namespace PlatformServices.Data;

public interface IPlatformRepo
{
    void CreatePlatform(Platform? plat);
    IEnumerable<Platform?> GetAllPlatforms();
    Platform? GetPlatformById(int id);
    bool SaveChanges();
}
