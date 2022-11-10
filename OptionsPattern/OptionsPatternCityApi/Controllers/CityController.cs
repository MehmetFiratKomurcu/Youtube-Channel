using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Options;

namespace OptionsPatternCityApi.Controllers;

[ApiController]
[Route("[controller]")]
public class CityController : ControllerBase
{
    private readonly IConfiguration _configuration;
    private readonly CityStatusOptions _optionsSnapshot;
    private readonly CityStatusOptions _optionsMonitor;
    private readonly CityStatusOptions _options;

    public CityController(IConfiguration configuration, IOptions<CityStatusOptions> options,
        IOptionsSnapshot<CityStatusOptions> optionsSnapshot, IOptionsMonitor<CityStatusOptions> optionsMonitor)
    {
        _configuration = configuration;
        _optionsSnapshot = optionsSnapshot.Value;
        _optionsMonitor = optionsMonitor.CurrentValue;
        _options = options.Value;
    }

    [HttpGet(Name = "CityStatus")]
    public object Get()
    {
        var name = _configuration.GetValue<string>("CityStatus:Name");
        var population = _configuration.GetValue<long>("CityStatus:Population");

        return new
        {
            Config = new { name, population },
            ConfigOptions = new { _options.Name, _options.Population },
            ConfigOptionsSnaphot = new { _optionsSnapshot.Name, _optionsSnapshot.Population },
            ConfigOptionsMonitor = new { _optionsMonitor.Name, _optionsMonitor.Population }
        };
    }
}