const https = require("https");

// token configuration
const clientId = "[[client_id]]";
const clientSecret = "[[client_secret]]";
const scopes = "[[scopes]]";

let token = "";

async function makeTokenRequest() {
  return new Promise((resolve, reject) => {
    const tokenPostData = JSON.stringify({
      client_id: clientId,
      client_secret: clientSecret,
      grant_type: "client_credentials",
      scope: scopes,
    });

    const tokenOptions = {
      hostname: "api.intelligent-api.com",
      port: 443,
      path: "/v1/token",
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Content-Length": Buffer.byteLength(tokenPostData),
      },
    };

    const tokenRequest = https.request(tokenOptions, (response) => {
      let responseBody = "";
      response.on("data", (chunk) => {
        responseBody += chunk;
      });

      response.on("end", () => {
        try {
          const jsonResponse = JSON.parse(responseBody);
          resolve(jsonResponse.access_token);
        } catch (error) {
          reject(`Error parsing JSON: ${error.message}`);
        }
      });
    });

    tokenRequest.on("error", (e) => {
      reject(`Problem with request: ${e.message}`);
    });

    tokenRequest.write(tokenPostData);
    tokenRequest.end();
  });
}

(async () => {
  try {
    token = await makeTokenRequest();

    // use the access_token as a bearer token
    const authorization = `Bearer ${token}`;

    // populate endpoint parameters
    const scrapeUrl = "[[url]]";
    const userAgent = "Intelligent API Sample Nodejs Code";

    const postData = JSON.stringify({
      url: scrapeUrl,
    });

    const options = {
      hostname: "api.intelligent-api.com",
      port: 443,
      path: "/v1/scraper/markdown",
      method: "POST",
      headers: {
        Authorization: authorization,
        "Content-Type": "application/json",
        "Content-Length": Buffer.byteLength(postData),
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

    request.write(postData);
    request.end();
  } catch (error) {
    console.error("Error:", error);
  }
})();
