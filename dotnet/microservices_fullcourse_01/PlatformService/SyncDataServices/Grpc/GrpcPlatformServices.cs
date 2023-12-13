using AutoMapper;
using Grpc.Core;
using PlatformServices.Data;

namespace PlatformService.SyncDataService.Grpc;

public class GrpcPlatformService : GrpcPlatform.GrpcPlatformBase
{
    private readonly IPlatformRepo _repository;
    private readonly IMapper _mapper;

    public GrpcPlatformService(IPlatformRepo repository, IMapper mapper)
    {
        _repository = repository;
        _mapper = mapper;
    }

    public override Task<GrpcPlatformResponse> GetAllPlatforms(
        GetAllRequest request,
        ServerCallContext context)
    {
        var response = new GrpcPlatformResponse();
        var platforms = _repository.GetAllPlatforms();

        foreach (var platform in platforms)
        {
            response.Platform.Add(_mapper.Map<GrpcPlatformModel>(platform));
        }

        return Task.FromResult(response);
    }
}