require 'base64'
require 'net/http'
require 'json'

# basic configuration
client_id = "[[client_id]]"
client_secret = "[[client_secret]]"

# create a base64 encoded api key
api_key = Base64.strict_encode64("#{client_id}:#{client_secret}")

# use the api key as a basic token
authorization = "Basic #{api_key}"

# populate endpoint parameters
text = "[[text_body]]"
source_language = "[[source_language]]"
target_language = "[[target_language]]"
user_agent = "Intelligent API Sample Ruby Code"

data = {
  text: text,
  sourceLanguage: source_language,
  targetLanguage: target_language
}

json_data = data.to_json

# invoke the API endpoint
url = URI("https://api.intelligent-api.com/v1/language/translate")

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