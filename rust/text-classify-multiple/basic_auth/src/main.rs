use base64::{engine::general_purpose::STANDARD, Engine as _};
use reqwest::header::{HeaderMap, HeaderValue, AUTHORIZATION, CONTENT_TYPE, USER_AGENT};
use reqwest::Client;
use serde::Serialize;

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
    let user_agent = "Intelligent API Sample Rust Code";

    #[derive(Serialize)]
    struct PostBody {
        items: Vec<String>,
        classifications: Vec<String>,
    }

    let post_body = PostBody {
        items: vec![
            "McDonalds Quarter Pounder deluxe".to_string(),
            "McDonalds Regular Coke".to_string(),
            "Gym Membership".to_string(),
            "Beef Steak".to_string(),
            "Multi-vitamin".to_string(),
            "Petrol".to_string(),
            "Car tires".to_string(),
        ],
        classifications: vec![
            "Food & Dining".to_string(),
            "Transport".to_string(),
            "Shopping".to_string(),
            "Utilities".to_string(),
            "Entertainment".to_string(),
            "Health".to_string(),
            "Education".to_string(),
            "Other".to_string(),
        ],
    };

    // invoke the API endpoint
    let url = "https://api.intelligent-api.com/v1/text/classify/multiple";

    let mut headers = HeaderMap::new();
    headers.insert(CONTENT_TYPE, HeaderValue::from_static("application/json"));
    headers.insert(AUTHORIZATION, HeaderValue::from_str(&authorization)?);
    headers.insert(USER_AGENT, HeaderValue::from_str(&user_agent)?);

    let client = Client::new();

    let response = client
        .post(url)
        .headers(headers)
        .json(&post_body)
        .send()
        .await?;

    println!("Response Status: {}", response.status());
    let body = response.text().await?;
    println!("Response Body: {}", body);

    Ok(())
}
