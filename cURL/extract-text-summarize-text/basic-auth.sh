# basic configuration
CLIENT_ID="[[client_id]]"
CLIENT_SECRET="[[client_secret]]"

# create a base64 encoded api key
API_KEY=$(echo -n "$CLIENT_ID:$CLIENT_SECRET" | base64)

# use the api key as a basic token
AUTHORIZATION="Basic $API_KEY"

# populate endpoint parameters
FULL_PATH_TO_FILE="[[filepath]]"
CONTENT_TYPE="[[content_type]]"

# invoke the API endpoint
curl --location 'https://api.intelligent-api.com/v1/document/text/summarize' \
--header "Content-Type: $CONTENT_TYPE" \
--header "Authorization: ${AUTHORIZATION}" \
--data-binary "@$FULL_PATH_TO_FILE"