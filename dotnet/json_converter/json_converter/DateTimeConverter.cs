using System.Globalization;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace json_converter;

public class DateTimeConverter : JsonConverter<DateTime>
{
    private readonly string _format;

    public DateTimeConverter(string format)
    {
        _format = format ?? throw new ArgumentNullException(nameof(format));
    }

    public override DateTime Read(
        ref Utf8JsonReader reader,
        Type typeToConvert,
        JsonSerializerOptions options)
    {
        if (reader.TokenType != JsonTokenType.String) throw new InvalidOperationException("Expected a string value for DateTime.");

        var dateString = reader.GetString();

        try
        {
            if (DateTime.TryParseExact(
                    dateString,
                    _format,
                    CultureInfo.InvariantCulture,
                    DateTimeStyles.None,
                    out var result))
            {
                return result;
            }
        }
        catch (FormatException ex)
        {
            throw new FormatException("Invalid date format.", ex);
        }

        throw new JsonException();
    }

    public override void Write(
        Utf8JsonWriter writer,
        DateTime value,
        JsonSerializerOptions options)
    {
        var token = value.ToString(_format, CultureInfo.InvariantCulture);
        writer.WriteStringValue(token ?? throw new ArgumentNullException(nameof(writer)));
    }
}
