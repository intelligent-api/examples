# token configuration
CLIENT_ID="[[client_id]]"
CLIENT_SECRET="[[client_secret]]"
SCOPES="[[scopes]]"

# using the jq plugin (https://jqlang.github.io/jq/download/) get the access_token from the token endpoint
ACCESS_TOKEN=$(curl --location 'https://api.intelligent-api.com/v1/token' \
--header 'Content-Type: application/json' \
--data "{ \"client_id\": \"$CLIENT_ID\", \"client_secret\": \"$CLIENT_SECRET\", \"grant_type\": \"client_credentials\", \"scope\": \"$SCOPES\" }" | jq -r '.access_token')

# use the access_token as a bearer token
AUTHORIZATION="Bearer $ACCESS_TOKEN"

# populate endpoint parameters
TEXT_BODY="[[text_body]]"

# invoke the API endpoint
curl --location 'https://api.intelligent-api.com/v1/language/detect/pii/redact' \
--header "Authorization: ${AUTHORIZATION}" \
--data "{ \"text\": \"$TEXT_BODY\" }"