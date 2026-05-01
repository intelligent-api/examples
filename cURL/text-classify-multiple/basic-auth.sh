# basic configuration
CLIENT_ID="[[client_id]]"
CLIENT_SECRET="[[client_secret]]"

# create a base64 encoded api key
API_KEY=$(echo -n "$CLIENT_ID:$CLIENT_SECRET" | base64)

# use the api key as a basic token
AUTHORIZATION="Basic $API_KEY"

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
