import java.io.ByteArrayOutputStream;
import java.io.File;
import java.io.OutputStreamWriter;
import java.io.PrintWriter;
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
                                        .POST(HttpRequest.BodyPublishers.ofString(tokenPayload))
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
                        String boundary = "----FormBoundary" + System.currentTimeMillis();

                        // Build multipart body
                        ByteArrayOutputStream bodyStream = new ByteArrayOutputStream();
                        PrintWriter writer = new PrintWriter(new OutputStreamWriter(bodyStream, "UTF-8"), true);

                        writer.append("--").append(boundary).append("\r\n");
                        writer.append("Content-Disposition: form-data; name=\"file\"; filename=\"")
                                        .append(file.getName())
                                        .append("\"").append("\r\n");
                        writer.append("Content-Type: application/octet-stream").append("\r\n");
                        writer.append("\r\n");
                        writer.flush();

                        bodyStream.write(Files.readAllBytes(file.toPath()));

                        writer.append("\r\n");
                        writer.append("--").append(boundary).append("--").append("\r\n");
                        writer.flush();

                        HttpClient client = HttpClient.newHttpClient();
                        HttpRequest request = HttpRequest.newBuilder()
                                        .uri(URI.create("https://api.intelligent-api.com/v1/speech/expenses?classifications=Food%20%26%20Dining,Transport,Shopping,Utilities,Entertainment,Health,Education,Other"))
                                        .header("Authorization", authorization)
                                        .header("User-Agent", userAgent)
                                        .header("Content-Type", "multipart/form-data; boundary=" + boundary)
                                        .POST(HttpRequest.BodyPublishers.ofByteArray(bodyStream.toByteArray()))
                                        .build();

                        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

                        System.out.println("Response Code: " + response.statusCode());
                        System.out.println("Response Body: " + response.body());
                } catch (Exception e) {
                        e.printStackTrace();
                }
        }
}
