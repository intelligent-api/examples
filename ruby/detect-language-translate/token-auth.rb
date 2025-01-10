require 'base64'
require 'net/http'
require 'json'

# token configuration
client_id = "[[client_id]]"
client_secret = "[[client_secret]]"
scopes = "[[scopes]]"

# get the access_token from the token endpoint
token_data = {
  client_id: client_id,
  client_secret: client_secret,
  grant_type: "client_credentials",
  scope: scopes
}

token_json_data = token_data.to_json
token_url = URI("https://api.intelligent-api.com/v1/token")

token_http = Net::HTTP.new(token_url.host, token_url.port)
token_http.use_ssl = (token_url.scheme == "https")

token_request = Net::HTTP::Post.new(token_url)
token_request["Content-Type"] = "application/json"
token_request.body = token_json_data

token_response = token_http.request(token_request)
token_response_data = JSON.parse(token_response.body)
token = token_response_data["access_token"]

# use the access_token as a bearer token
authorization = "Bearer #{token}"

# populate endpoint parameters
text = "[[text_body]]"
target_language = "[[target_language]]"
user_agent = "Intelligent API Sample Ruby Code"

data = {
  text: text,
  targetLanguage: target_language
}

json_data = data.to_json

# invoke the API endpoint
url = URI("https://api.intelligent-api.com/v1/language/detect/translate")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = (url.scheme == "https")

request = Net::HTTP::Post.new(url)
request["Authorization"] = authorization
request["Content-Type"] = "application/json"
request["User-Agent"] = user_agent
request.body = json_data

response = http.request(request)

puts "Response Code: #{response.code}"
puts "Response Body: #{response.body}"