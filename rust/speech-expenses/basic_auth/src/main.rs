use base64::{engine::general_purpose::STANDARD, Engine as _};
use reqwest::header::{HeaderValue, AUTHORIZATION, USER_AGENT};
use reqwest::Client;
use reqwest::multipart::{Form, Part};
use std::fs::File;
use std::io::{Read};
use tokio::fs::File;
use tokio::io::AsyncReadExt;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // basic configuration
    let client_id: &str = "[[client_id]]";
    let client_secret: &str = "[[client_secret]]";

    // create a base64 encoded api key
    let api_key = STANDARD.encode(format!("{client_id}:{client_secret}"));

    // use the api key as a basic token
    let authorization = format!("Basic {api_key}");

    // populate endpoint parameters
    let file_path = "[[filepath]]";
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
    let url = "https://api.intelligent-api.com/v1/speech/expenses?classifications=Food%20%26%20Dining,Transport,Shopping,Utilities,Entertainment,Health,Education,Other";

    let client = Client::new();

    let response = client.post(url)
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
