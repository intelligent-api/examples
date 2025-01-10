import base64
import http.client
import json

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

with open(file_path, "rb") as file:
    binary_data = file.read()

# invoke the API endpoint
host = "api.intelligent-api.com"
path = "/v1/document/expenses"

connection = http.client.HTTPSConnection(host)

headers = {
    "Authorization": authorization,
    "Content-Type": "application/octet-stream",
    "User-Agent": user_agent
}

connection.request("POST", path, body=binary_data, headers=headers)
response = connection.getresponse()
response_data = response.read().decode("utf-8")

print("Status Code:", response.status)
print("Response:", response_data)

connection.close()