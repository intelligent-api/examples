# basic configuration
CLIENT_ID="[[client_id]]"
CLIENT_SECRET="[[client_secret]]"

# create a base64 encoded api key
API_KEY=$(echo -n "$CLIENT_ID:$CLIENT_SECRET" | base64)

# use the api key as a basic token
AUTHORIZATION="Basic $API_KEY"

# populate endpoint parameters
TEXT_BODY="[[text_body]]"
SOURCE_LANGUAGE="[[source_language]]"
TARGET_LANGUAGE="[[target_language]]"

# invoke the API endpoint
curl --location 'https://api.intelligent-api.com/v1/language/translate' \
--header "Authorization: ${AUTHORIZATION}" \
--data "{ \"text\": \"$TEXT_BODY\", \"sourceLanguage\": \"$SOURCE_LANGUAGE\", \"targetLanguage\": \"$TARGET_LANGUAGE\" }"