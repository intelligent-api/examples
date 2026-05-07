import base64
import http.client
import mimetypes
import os
import uuid

# basic configuration
client_id = "[[client_id]]"
client_secret = "[[client_secret]]"

# create a base64 encoded api key
api_key_bytes = base64.b64encode((client_id + ":" + client_secret).encode("utf-8"))
api_key = api_key_bytes.decode("utf-8")

# use the api key as a basic token
authorization = "Basic " + api_key

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
path = "/v1/speech/expenses?classifications=Food%20%26%20Dining,Transport,Shopping,Utilities,Entertainment,Health,Education,Other"

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
