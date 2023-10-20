using Microsoft.AspNetCore.Mvc;

namespace CommandsService.Controllers;

[Route("api/c/[controller]")]
[ApiController]
public class PlatformsController : ControllerBase
{
    public PlatformsController()
    {

    }

    public ActionResult TestInboundConnection()
    {
        Console.WriteLine("--> Inbound test of from Platforms # Command Service");

        return Ok("Inbound test of from Platforms Controller");
    }
}