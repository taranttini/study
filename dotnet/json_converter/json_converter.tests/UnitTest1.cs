namespace json_converter.tests;

using json_converter;
using Xunit;

public class UnitTest1
{
    // SUT = system under test
    private readonly DateTimeConverter _sut = new("yyyy-MM-dd H:mm");

    [Fact]
    public void Can_read_string_value_as_datetime()
    {
        var result = _sut.Read("\"2023-08-01 6:00\"");
        Assert.Equal(new DateTime(2023, 8, 1, 6, 0, 0), result);
    }

    [Fact]
    public void Can_write_datetime_as_string()
    {
        var result = _sut.Write(new DateTime(2023, 8, 1, 6, 0, 0));
        Assert.Equal("\"2023-08-01 6:00\"", result);
    }

    [Fact]
    public void Can_read_string_value_as_datetime_using_quote()
    {
        var result = _sut.Read((Quote)"2023-08-01 6:00");
        Assert.Equal(new DateTime(2023, 8, 1, 6, 0, 0), result);
    }
}