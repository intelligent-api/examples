// token configuration
var clientId = "[[client_id]]";
var clientSecret = "[[client_secret]]";
var scopes = "[[scopes]]";

dynamic tokenRequest = new
{
    client_id = clientId,
    client_secret = clientSecret,
    grant_type = "client_credentials",
    scope = scopes
};

var tokenRequestJson = System.Text.Json.JsonSerializer.Serialize(tokenRequest);
var authorization = string.Empty;

using (var tokenClient = new HttpClient())
{
    var request = new HttpRequestMessage(HttpMethod.Post, "https://api.intelligent-api.com/v1/token");
    var content = new StringContent(tokenRequestJson, null, "application/json");
    request.Content = content;
    var response = await tokenClient.SendAsync(request);
    response.EnsureSuccessStatusCode();
    var responseString = await response.Content.ReadAsStringAsync();
    dynamic responseObject = System.Text.Json.JsonSerializer.Deserialize<System.Dynamic.ExpandoObject>(responseString);
    authorization = $"Bearer {responseObject.access_token}";
}

// populate endpoint paramaters
var scrapeUrl = "[[url]]";
var userAgent = "Intelligent API Sample C# Code";

dynamic requestBody = new
{
    url = scrapeUrl
};

var requestBodyJson = System.Text.Json.JsonSerializer.Serialize(requestBody);

// invoke the api endpoint
using (var client = new HttpClient())
{
    var request = new HttpRequestMessage(HttpMethod.Post, "https://api.intelligent-api.com/v1/scraper/markdown");
    request.Headers.Add("Authorization", authorization);
    request.Headers.Add("User-Agent", userAgent);
    var content = new StringContent(requestBodyJson, null, "application/json");
    request.Content = content;
    var response = await client.SendAsync(request);
    response.EnsureSuccessStatusCode();
    Console.WriteLine(await response.Content.ReadAsStringAsync());
}
