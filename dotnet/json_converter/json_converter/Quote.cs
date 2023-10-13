namespace json_converter;

public class Quote  
{  
    private readonly object _value;  
    private Quote(object value) => _value = value;  
  
    public static explicit operator Quote(string value) => new(value);  
  
    public static implicit operator string(Quote value) => value.ToString();  
    public override string ToString() => $"\"{_value}\"";  
}
