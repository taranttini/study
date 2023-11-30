using AutoMapper;
using CommandsService.Dto;
using CommandsService.Dtos;
using CommandsService.Models;
using Microsoft.AspNetCore.Mvc;

namespace CommandsService.Controllers;

[Route("api/c/platforms/{platformId}/[controller]")]
[ApiController]
public class CommandsController : ControllerBase
{
    private readonly ICommandRepo _commandRepo;
    private readonly IMapper _mapper;

    public CommandsController(ICommandRepo commandRepo, IMapper mapper)
    {
        _commandRepo = commandRepo;
        _mapper = mapper;
    }

    [HttpGet]
    public ActionResult<IEnumerable<CommandReadDto>> GetCommandsForPlatforms(int platformId)
    {
        Console.WriteLine($"--> Hit GetComandsForPlatforms: {platformId}");

        if (!_commandRepo.PlaftormExists(platformId))
        {
            return NotFound();
        }

        var commands = _commandRepo.GetCommandsForPlatform(platformId);

        return Ok(_mapper.Map<IEnumerable<CommandReadDto>>(commands));
    }

    [HttpGet("{commandId}", Name = "GetCommandForPlatform")]
    public ActionResult<IEnumerable<CommandReadDto>> GetCommandForPlatform(int platformId, int commandId)
    {
        Console.WriteLine($"--> Hit GetCommandForPlatform {platformId} / {commandId}");

        if (!_commandRepo.PlaftormExists(platformId))
        {
            return NotFound();
        }

        var commands = _commandRepo.GetCommand(platformId, commandId);

        if (commands is null)
        {
            return NotFound();
        }

        return Ok(_mapper.Map<IEnumerable<CommandReadDto>>(commands));
    }

    [HttpPost] //("{commandId}", Name = "CreateCommandForPlatform")]
    public ActionResult<IEnumerable<CommandReadDto>> CreateCommandForPlatform(int platformId, CommandCreateDto commandDto)
    {
        Console.WriteLine($"--> Hit CreateCommandForPlatform {platformId}");

        if (!_commandRepo.PlaftormExists(platformId))
        {
            return NotFound();
        }

        var command = _mapper.Map<Command>(commandDto);

        _commandRepo.CreateCommand(platformId, command);
        _commandRepo.SaveChanges();


        var commandReadDto = _mapper.Map<CommandReadDto>(command);

        return CreatedAtRoute(nameof(GetCommandForPlatform), new
        { platformId, commandId = commandReadDto.Id }, commandReadDto);
    }

}