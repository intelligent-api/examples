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
users_date = "[[users_date]]"
users_day = "[[users_day]]"
users_time = "[[users_time]]"
user_agent = "Intelligent API Sample Ruby Code"

# build multipart form body
boundary = "----FormBoundary#{SecureRandom.hex}"
file_data = File.read(file_path, mode: "rb")
filename = File.basename(file_path)
mime_type = MIME::Types.type_for(file_path).first&.content_type || "application/octet-stream"

body = []
body << "--#{boundary}\r\n"
body << "Content-Disposition: form-data; name=\"file\"; filename=\"#{filename}\"\r\n"
body << "Content-Type: #{mime_type}\r\n"
body << "\r\n"
body << file_data
body << "\r\n--#{boundary}--\r\n"

multipart_body = body.join

# invoke the API endpoint
url = URI("https://api.intelligent-api.com/v1/speech/todo?user_date=#{users_date}&user_day=#{users_day}&user_time=#{users_time}")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = (url.scheme == "https")

request = Net::HTTP::Post.new(url)
request["Authorization"] = authorization
request["Content-Type"] = "multipart/form-data; boundary=#{boundary}"
request["User-Agent"] = user_agent
request.body = multipart_body

response = http.request(request)

puts "Response Code: #{response.code}"
puts "Response Body: #{response.body}"
