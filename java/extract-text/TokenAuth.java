import java.io.File;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.file.Files;

public class TokenAuth {
        public static void main(String args[]) {
                // token configuration
                String clientId = "[[client_id]]";
                String clientSecret = "[[client_secret]]";
                String scopes = "[[scopes]]";

                // get the access_token from the token endpoint
                String tokenPayload = "{ \"client_id\": \"" + clientId + "\", \"client_secret\": \"" + clientSecret
                                + "\", \"grant_type\": \"client_credentials\", \"scope\": \"" + scopes + "\"  }";
                String token = "";

                try {
                        HttpClient client = HttpClient.newHttpClient();
                        HttpRequest request = HttpRequest.newBuilder()
                                        .uri(URI.create("https://api.intelligent-api.com/v1/token"))
                                        .header("Content-Type", "application/json")
                                        .POST(HttpRequest.BodyPublishers.ofString(
                                                        tokenPayload))
                                        .build();

                        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

                        String searchKey = "access_token";
                        String tokenResponse = response.body();
                        int keyIndex = tokenResponse.indexOf(searchKey);
                        int valueStart = tokenResponse.indexOf("\"", keyIndex + searchKey.length()) + 3;
                        int valueEnd = tokenResponse.indexOf("\"", valueStart);
                        token = tokenResponse.substring(valueStart, valueEnd);
                } catch (Exception e) {
                        e.printStackTrace();
                        return;
                }

                // use the access_token as a bearer token
                String authorization = "Bearer ".concat(token);

                // populate endpoint parameters
                String fullPathToFile = "[[filepath]]";
                String userAgent = "Intelligent API Sample Java Code";

                // invoke the API endpoint
                try {
                        File file = new File(fullPathToFile);
                        byte[] fileBytes = Files.readAllBytes(file.toPath());
                        HttpClient client = HttpClient.newHttpClient();
                        HttpRequest request = HttpRequest.newBuilder()
                                        .uri(URI.create("https://api.intelligent-api.com/v1/document/text"))
                                        .header("Authorization",
                                                        authorization)
                                        .header("User-Agent",
                                                        userAgent)
                                        .header("Content-Type", "application/octet-stream")
                                        .POST(HttpRequest.BodyPublishers.ofByteArray(
                                                        fileBytes))
                                        .build();

                        HttpResponse<String> response = client.send(request,
                                        HttpResponse.BodyHandlers.ofString());

                        System.out.println("Response Code: " + response.statusCode());
                        System.out.println("Response Body: " + response.body());
                } catch (Exception e) {
                        e.printStackTrace();
                }
        }
}
