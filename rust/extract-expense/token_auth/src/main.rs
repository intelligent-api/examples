use reqwest::header::{HeaderMap, HeaderValue, CONTENT_TYPE, AUTHORIZATION, USER_AGENT};
use reqwest::Client;
use serde::{Deserialize, Serialize};
use std::fs::File;
use std::io::{Read};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // token configuration
    let client_id: &str = "[[client_id]]";
    let client_secret: &str = "[[client_secret]]";
    let scopes: &str = "[[scopes]]";

    // get the access_token from the token endpoint
    #[derive(Serialize)]
    struct TokenPostBody {
        client_id: String,
        client_secret: String,
        grant_type: String,
        scope: String,
    }

    #[derive(Deserialize)]
    struct TokenApiResponse {
        access_token: String,
    }

    let token_post_body = TokenPostBody {
        client_id: client_id.to_string(),
        client_secret: client_secret.to_string(),
        grant_type: "client_credentials".to_string(),
        scope: scopes.to_string(),
    };

    let token_url = "https://api.intelligent-api.com/v1/token";

    let mut token_headers = HeaderMap::new();
    token_headers.insert(CONTENT_TYPE, HeaderValue::from_static("application/json"));

    let token_client = Client::new();

    let token_response = token_client.post(token_url)
        .headers(token_headers)
        .json(&token_post_body)
        .send()
        .await?
        .json::<TokenApiResponse>()
        .await?;

    let token = token_response.access_token.to_string();

    // use the access_token as a bearer token
    let authorization = format!("Bearer {token}");

    // populate endpoint parameters
    let file_path = "[[filepath]]";
    let user_agent = "Intelligent API Sample Rust Code";

    let mut file = File::open(file_path)?;
    let mut file_data = Vec::new();
    file.read_to_end(&mut file_data)?;

    // invoke the API endpoint
    let url = "https://api.intelligent-api.com/v1/document/expenses";

    let mut headers = HeaderMap::new();
    headers.insert(CONTENT_TYPE, HeaderValue::from_static("application/octet-stream"));
    headers.insert(AUTHORIZATION, HeaderValue::from_str(&authorization)?);
    headers.insert(USER_AGENT, HeaderValue::from_str(&user_agent)?);

    let client = Client::new();

    let response = client.post(url)
        .headers(headers)
        .body(file_data)
        .send()
        .await?;

    println!("Response Status: {}", response.status());
    let body = response.text().await?;
    println!("Response Body: {}", body);

    Ok(())
}
