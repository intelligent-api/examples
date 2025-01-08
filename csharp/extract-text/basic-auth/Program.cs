// basic configuration
var clientId = "[[client_id]]";
var clientSecret = "[[client_secret]]";

// create a base64 encoded api key
var apiKeyText = $"{clientId}:{clientSecret}";
var apiKey = Convert.ToBase64String(System.Text.Encoding.UTF8.GetBytes(apiKeyText));

// use the api key as the basic token
var authorization = $"Basic {apiKey}";

// populate endpoint paramaters
var fullPathToFile = "[[filepath]]";
var userAgent = "Intelligent API Sample C# Code";

// invoke the api endpoint
using (var client = new HttpClient())
{
    var request = new HttpRequestMessage(HttpMethod.Post, "https://api.intelligent-api.com/v1/document/text");
    request.Headers.Add("Authorization", authorization);
    request.Headers.Add("User-Agent", userAgent);
    request.Content = new StreamContent(File.OpenRead(fullPathToFile));
    var response = await client.SendAsync(request);
    response.EnsureSuccessStatusCode();
    Console.WriteLine(await response.Content.ReadAsStringAsync());
}
