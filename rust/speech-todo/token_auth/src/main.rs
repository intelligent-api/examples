use reqwest::header::{HeaderValue, AUTHORIZATION, USER_AGENT};
use reqwest::Client;
use reqwest::multipart::{Form, Part};
use serde::{Deserialize, Serialize};
use std::fs::File;
use std::io::{Read};
use tokio::fs::File;
use tokio::io::AsyncReadExt;

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

    let token_client = Client::new();

    let token_response = token_client.post(token_url)
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
    let users_date = "[[users_date]]";
    let users_day = "[[users_day]]";
    let users_time = "[[users_time]]";
    let user_agent = "Intelligent API Sample Rust Code";

    let mut file = File::open(file_path).await?;
    let mut file_data = Vec::new();
    file.read_to_end(&mut file_data).await?;

    let file_name = Path::new(file_path)
        .file_name()
        .and_then(|n| n.to_str())
        .unwrap_or("file")
        .to_string();

    // build multipart form
    let part = Part::bytes(file_data)
        .file_name(file_name)
        .mime_str("application/octet-stream")?;

    let form = Form::new().part("file", part);

    // invoke the API endpoint
    let url = format!("https://api.intelligent-api.com/v1/speech/todo?user_date={users_date}&user_day={users_day}&user_time={users_time}");

    let client = Client::new();

    let response = client.post(&url)
        .header(AUTHORIZATION, HeaderValue::from_str(&authorization)?)
        .header(USER_AGENT, HeaderValue::from_str(&user_agent)?)
        .multipart(form)
        .send()
        .await?;

    println!("Response Status: {}", response.status());
    let body = response.text().await?;
    println!("Response Body: {}", body);

    Ok(())
}
