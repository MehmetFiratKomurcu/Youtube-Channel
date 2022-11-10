using System.ComponentModel.DataAnnotations;

namespace OptionsPatternCityApi;

public class CityStatusOptions
{
    public const string CityStatus = "CityStatus";
    public string Name { get; set; } = string.Empty;
    [Range(0, long.MaxValue)]
    public long Population { get; set; } = 0;
}