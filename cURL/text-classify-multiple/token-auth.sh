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

# invoke the API endpoint
curl --location 'https://api.intelligent-api.com/v1/text/classify/multiple' \
--header "Authorization: ${AUTHORIZATION}" \
--header 'Content-Type: application/json' \
--data '{
  "items": [
    "McDonalds Quarter Pounder deluxe",
    "McDonalds Regular Coke",
    "Gym Membership",
    "Beef Steak",
    "Multi-vitamin",
    "Petrol",
    "Car tires"
  ],
  "classifications": [
    "Food & Dining",
    "Transport",
    "Shopping",
    "Utilities",
    "Entertainment",
    "Health",
    "Education",
    "Other"
  ]
}'
