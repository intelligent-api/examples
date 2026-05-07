import java.io.ByteArrayOutputStream;
import java.io.File;
import java.io.OutputStreamWriter;
import java.io.PrintWriter;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.file.Files;
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
