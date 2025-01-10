const https = require("https");

// basic configuration
const clientId = "[[client_id]]";
const clientSecret = "[[client_secret]]";

// create a base64 encoded api key
const apiKey = Buffer.from(`${clientId}:${clientSecret}`).toString("base64");

(async () => {
  try {
    // use the api key as a basic token
    const authorization = `Basic ${apiKey}`;

    // populate endpoint parameters
    const fs = require("fs");
    const fullPathToFile = "[[filepath]]";
    const userAgent = "Intelligent API Sample Nodejs Code";

    const fileStream = fs.createReadStream(fullPathToFile);

    const options = {
      hostname: "api.intelligent-api.com",
      port: 443,
      path: "/v1/document/text",
      method: "POST",
      headers: {
        Authorization: authorization,
        "Content-Type": "application/octet-stream",
        "User-Agent": userAgent,
      },
    };

    const request = https.request(options, (response) => {
      console.log(`Status Code: ${response.statusCode}`);

      response.on("data", (chunk) => {
        console.log("Response Body:", chunk.toString());
      });
    });

    request.on("error", (e) => {
      console.error(`Problem with request: ${e.message}`);
    });

    fileStream.pipe(request);
    fileStream.on("end", () => {
      request.end();
    });
  } catch (error) {
    console.error("Error:", error);
  }
})();
