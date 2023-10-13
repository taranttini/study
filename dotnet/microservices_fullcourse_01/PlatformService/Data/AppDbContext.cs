using Microsoft.EntityFrameworkCore;
using PlatformService.Models;

namespace PlatformService.Data;

public class AppDbContext : DbContext
{
    public AppDbContext(DbContextOptions<AppDbContext> opt, DbSet<Platform> platforms) : base(opt)
    {
        Platforms = platforms;
    }

    public DbSet<Platform> Platforms { get; set; }
}
