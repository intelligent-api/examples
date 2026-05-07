use base64::{engine::general_purpose::STANDARD, Engine as _};
use reqwest::header::{HeaderMap, HeaderValue, CONTENT_TYPE, AUTHORIZATION, USER_AGENT};
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
    let text: &str = "[[text_body]]";
    let users_date: &str = "[[users_date]]";
    let users_day: &str = "[[users_day]]";
    let users_time: &str = "[[users_time]]";
    let user_agent = "Intelligent API Sample Rust Code";

    #[derive(Serialize)]
    struct PostBody {
        text: String,
        #[serde(rename = "usersDate")]
        users_date: String,
        #[serde(rename = "usersDay")]
        users_day: String,
        #[serde(rename = "usersTime")]
        users_time: String,
    }

    let post_body = PostBody {
        text: text.to_string(),
        users_date: users_date.to_string(),
        users_day: users_day.to_string(),
        users_time: users_time.to_string(),
    };

    // invoke the API endpoint
    let url = "https://api.intelligent-api.com/v1/text/todo";

    let mut headers = HeaderMap::new();
    headers.insert(CONTENT_TYPE, HeaderValue::from_static("application/json"));
    headers.insert(AUTHORIZATION, HeaderValue::from_str(&authorization)?);
    headers.insert(USER_AGENT, HeaderValue::from_str(&user_agent)?);

    let client = Client::new();

    let response = client.post(url)
        .headers(headers)
        .json(&post_body)
        .send()
        .await?;

    println!("Response Status: {}", response.status());
    let body = response.text().await?;
    println!("Response Body: {}", body);

    Ok(())
}
