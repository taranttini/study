﻿using CommandsService.Models;

namespace CommandsService;

public class CommandRepo : ICommandRepo
{
    private readonly AppDbContext _context;

    public CommandRepo(AppDbContext context)
    {
        _context = context;
    }
    public void CreateCommand(int platformId, Command command)
    {
        if (command == null)
        {
            throw new ArgumentNullException(nameof(command));

        }
        command.PlatformId = platformId;
        _context.Commands.Add(command);
    }

    public void CreatePlatform(Platform platform)
    {
        if (platform == null)
        {
            throw new ArgumentNullException(nameof(platform));
        }

        _context.Platforms.Add(platform);
    }

    public bool ExternalPlatormExist(int externalPlataformId)
    {
        return _context.Platforms.Any(q => q.ExternalID == externalPlataformId);
    }

    public IEnumerable<Platform> GetAllPlatforms()
    {
        return _context.Platforms.ToList();
    }

    public Command? GetCommand(int platformId, int commandId)
    {
        return _context
            .Commands
            .Where(c => c.PlatformId == platformId && c.Id == commandId)
            .FirstOrDefault();
    }

    public IEnumerable<Command> GetCommandsForPlatform(int platformId)
    {
        return _context.Commands
            .Where(c => c.PlatformId == platformId)
            .OrderBy(c => c.Platform!.Name);
    }

    public bool PlaftormExists(int platformId)
    {
        return _context.Platforms.Any(q => q.Id == platformId);
    }

    public bool SaveChanges()
    {
        return _context.SaveChanges() >= 0;
    }
}
