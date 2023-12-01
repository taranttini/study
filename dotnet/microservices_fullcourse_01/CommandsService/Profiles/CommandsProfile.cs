using AutoMapper;
using CommandsService.Dto;
using CommandsService.Dtos;
using CommandsService.Models;

namespace CommandsProfile.Profiles;


public class CommandsProfile : Profile
{
    public CommandsProfile()
    {
        CreateMap<Platform, PlatformReadDto>();
        CreateMap<CommandCreateDto, Command>();
        CreateMap<Command, CommandReadDto>();
        CreateMap<PlatformPublishedDto, Platform>()
            .ForMember(
                destination => destination.ExternalID,
                options => options.MapFrom(source => source.Id)
            );
    }
}