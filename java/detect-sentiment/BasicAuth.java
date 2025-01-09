import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.Base64;

public class BasicAuth {
    public static void main(String args[]) {
        // basic configuration
        String clientId = "[[client_id]]";
        String clientSecret = "[[client_secret]]";

        // create a base64 encoded api key
        String apiKey = Base64.getEncoder().encodeToString((clientId.concat(":")).concat(clientSecret).getBytes());

        // use the api key as a basic token
        String authorization = "Basic ".concat(apiKey);

        // populate endpoint parameters
        String textBody = "[[text_body]]";
        String language = "[[source_language]]";
        String userAgent = "Intelligent API Sample Java Code";

        String jsonPayload = "{ \"text\": \"" + textBody + "\", \"language\": \"" + language + "\" }";

        // invoke the API endpoint
        try {
            HttpClient client = HttpClient.newHttpClient();
            HttpRequest request = HttpRequest.newBuilder()
                    .uri(URI.create("https://api.intelligent-api.com/v1/text/sentiment"))
                    .header("Authorization",
                            authorization)
                    .header("User-Agent",
                            userAgent)
                    .header("Content-Type", "application/json")
                    .POST(HttpRequest.BodyPublishers.ofString(
                            jsonPayload))
                    .build();

            HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

            System.out.println("Response Code: " + response.statusCode());
            System.out.println("Response Body: " + response.body());
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
