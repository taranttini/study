using AutoMapper;
using CommandsService.Dto;
using CommandsService.Dtos;
using CommandsService.Models;
using PlatformService;

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
        CreateMap<GrpcPlatformModel, Platform>()
        .ForMember(dest => dest.ExternalID, opt => opt.MapFrom(src => src.PlatformId))
        .ForMember(dest => dest.Name, opt => opt.MapFrom(src => src.Name))
        .ForMember(dest => dest.Commands, opt => opt.Ignore());
    }
}