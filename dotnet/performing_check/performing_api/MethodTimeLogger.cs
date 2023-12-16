using System.Reflection;

namespace performing_api;

public class MethodTimeLogger
{
    public static ILogger Logger;
    
    public static void Log(MethodBase methodBase, TimeSpan timeSpan, string message)
    {
        Logger.LogTrace("XXX {Class}.{Method} - {Message} in {Duration}",
            methodBase.DeclaringType!.Name, methodBase.Name, message, timeSpan);
        //Console.WriteLine();
    }
}