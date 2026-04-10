import http.client
import json
import mimetypes
import os
import uuid

# token configuration
client_id = "[[client_id]]"
client_secret = "[[client_secret]]"
scopes = "[[scopes]]"

# get the access_token from the token endpoint
token_data = {
    "client_id": client_id,
    "client_secret": client_secret,
    "grant_type": "client_credentials",
    "scope": scopes,
}

token_json_data = json.dumps(token_data)

token_host = "api.intelligent-api.com"
token_path = "/v1/token"

token_connection = http.client.HTTPSConnection(token_host)

token_headers = {"Content-Type": "application/json"}

token_connection.request(
    "POST", token_path, body=token_json_data, headers=token_headers
)
token_response = token_connection.getresponse()
token_response_data = token_response.read().decode("utf-8")

token_json_response = json.loads(token_response_data)
token = token_json_response.get("access_token")

# use the access_token as a bearer token
authorization = "Bearer " + token

# populate endpoint parameters
file_path = "[[filepath]]"
user_agent = "Intelligent API Sample Python Code"

# build multipart form body
boundary = uuid.uuid4().hex
body = []

# form part headers
body.append(f"--{boundary}".encode())
body.append(
    f'Content-Disposition: form-data; name="file"; filename="{os.path.basename(file_path)}"'.encode()
)
mime_type = mimetypes.guess_type(file_path)[0] or "application/octet-stream"
body.append(f"Content-Type: {mime_type}".encode())
body.append(b"")

# file content
with open(file_path, "rb") as file:
    body.append(file.read())

# closing boundary
body.append(f"--{boundary}--".encode())

multipart_body = b"\r\n".join(body)

# invoke the API endpoint
host = "api.intelligent-api.com"
path = "/v1/speech/transcribe"

connection = http.client.HTTPSConnection(host)

headers = {
    "Authorization": authorization,
    "Content-Type": f"multipart/form-data; boundary={boundary}",
    "User-Agent": user_agent,
}

connection.request("POST", path, body=multipart_body, headers=headers)
response = connection.getresponse()
response_data = response.read().decode("utf-8")

print("Status Code:", response.status)
print("Response:", response_data)

connection.close()
