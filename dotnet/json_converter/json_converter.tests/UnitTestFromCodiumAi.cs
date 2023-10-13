using System.Text;
using System.Text.Json;

namespace json_converter.tests;

using json_converter;
using Xunit;

public class UnitTestFromCodiumAi
{
    [Fact]
    public void test_valid_date_string_to_datetime_object()
    {
        // Arrange
        var converter = new DateTimeConverter("yyyy-MM-dd");
        var jsonString = "\"2022-01-01\"";
        var reader = new Utf8JsonReader(Encoding.UTF8.GetBytes(jsonString));
        var options = new JsonSerializerOptions();

        // Act
        var result = converter.Read(ref reader, typeof(DateTime), options);

        // Assert
        Assert.Equal(new DateTime(2022, 1, 1), result);
    }

    /*[Fact]
    public void test_datetime_object_to_string_with_format()
    {
        // Arrange
        var converter = new DateTimeConverter("yyyy-MM-dd");
        var dateTime = new DateTime(2022, 1, 1);
        var writer = new Utf8JsonWriter(Stream.Null);
        var options = new JsonSerializerOptions();

        // Act
        converter.Write(writer, dateTime, options);
        writer.Flush();
        var jsonString = Encoding.UTF8.GetString(writer.GetWrittenSpan());

        // Assert
        Assert.Equal("\"2022-01-01\"", jsonString);
    }*/

    [Fact]
    public void test_null_value_for_datetime_object()
    {
        // Arrange
        var converter = new DateTimeConverter("yyyy-MM-dd");
        var jsonString = "null";
        var reader = new Utf8JsonReader(Encoding.UTF8.GetBytes(jsonString));
        var options = new JsonSerializerOptions();

        // Act
        var result = converter.Read(ref reader, typeof(DateTime), options);

        // Assert
        Assert.Equal(default(DateTime), result);
    }

    [Fact]
    public void test_format_exception_for_invalid_date_format()
    {
        // Arrange
        var converter = new DateTimeConverter("yyyy-MM-dd");
        var jsonString = "\"2022/01/01\"";
        var options = new JsonSerializerOptions();

        // Act & Assert

        Assert.Throws<FormatException>(() =>
        {
            var reader = new Utf8JsonReader(Encoding.UTF8.GetBytes(jsonString));
            return converter.Read(ref reader, typeof(DateTime), options);
        });
    }

    [Fact]
    public void test_invalid_operation_exception_for_non_string_token_type()
    {
        // Arrange
        var converter = new DateTimeConverter("yyyy-MM-dd");
        var jsonString = "42";
        var reader = new Utf8JsonReader(Encoding.UTF8.GetBytes(jsonString));
        var options = new JsonSerializerOptions();

        // Act & Assert
        var action = converter.Read(ref reader, typeof(DateTime), options);
        Assert.Throws<InvalidOperationException>(() => action);
    }

    [Fact]
    public void test_json_exception_for_unparsable_date_string()
    {
        // Arrange
        var converter = new DateTimeConverter("yyyy-MM-dd");
        var jsonString = "\"2022-01-32\"";
        var reader = new Utf8JsonReader(Encoding.UTF8.GetBytes(jsonString));
        var options = new JsonSerializerOptions();

        // Act & Assert
        var action = converter.Read(ref reader, typeof(DateTime), options);
        Assert.Throws<JsonException>(() => action);
    }
}