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
file_path = "[[filepath]]"
user_agent = "Intelligent API Sample Ruby Code"

file_data = File.read(file_path, mode: "rb")

# invoke the API endpoint
url = URI("https://api.intelligent-api.com/v1/document/text/summarize")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = (url.scheme == "https")

request = Net::HTTP::Post.new(url)
request["Authorization"] = authorization
request["Content-Type"] = "application/octet-stream"
request["User-Agent"] = user_agent
request.body = file_data

response = http.request(request)

puts "Response Code: #{response.code}"
puts "Response Body: #{response.body}"