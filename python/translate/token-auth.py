import base64
import http.client
import json

# basic configuration
client_id = "[[client_id]]"
client_secret = "[[client_secret]]"
scopes = "[[scopes]]"

# get the access_token from the token endpoint
token_data = {
    "client_id": client_id,
    "client_secret": client_secret,
    "grant_type": "client_credentials",
    "scope": scopes
}

token_json_data = json.dumps(token_data)

token_host = "api.intelligent-api.com"
token_path = "/v1/token"

token_connection = http.client.HTTPSConnection(token_host)

token_headers = {
    "Content-Type": "application/json"
}

token_connection.request("POST", token_path, body=token_json_data, headers=token_headers)
token_response = token_connection.getresponse()
token_response_data = token_response.read().decode("utf-8")

token_json_response = json.loads(token_response_data)
token = token_json_response.get("access_token") 

# use the access_token as a bearer token
authorization = "Bearer " + token

# populate endpoint parameters
text = "[[text_body]]"
source_language = "[[source_language]]"
target_language = "[[target_language]]"
user_agent = "Intelligent API Sample Python Code"

data = {
    "text": text,
    "sourceLanguage": source_language,
    "targetLanguage": target_language
}

json_data = json.dumps(data)

# invoke the API endpoint
host = "api.intelligent-api.com"
path = "/v1/language/translate"

connection = http.client.HTTPSConnection(host)

headers = {
    "Authorization": authorization,
    "Content-Type": "application/json",
    "User-Agent": user_agent
}

connection.request("POST", path, body=json_data, headers=headers)
response = connection.getresponse()
response_data = response.read().decode("utf-8")

print("Status Code:", response.status)
print("Response:", response_data)

connection.close()