package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// basic configuration
	clientId := "[[client_id]]"
	clientSecret := "[[client_secret]]"

	// create a base64 encoded api key
	apiKey := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, clientSecret)))

	// use the api key as a basic token
	authorization := fmt.Sprintf("Basic %s", apiKey)

	// populate endpoint parameters
	userAgent := "Intelligent API Sample Go Code"

	type PostBody struct {
		Items           []string `json:"items"`
		Classifications []string `json:"classifications"`
	}

	postBody := PostBody{
		Items: []string{
			"McDonalds Quarter Pounder deluxe",
			"McDonalds Regular Coke",
			"Gym Membership",
			"Beef Steak",
			"Multi-vitamin",
			"Petrol",
			"Car tires",
		},
		Classifications: []string{
			"Food & Dining",
			"Transport",
			"Shopping",
			"Utilities",
			"Entertainment",
			"Health",
			"Education",
			"Other",
		},
	}

	payload, err := json.Marshal(postBody)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return
	}

	// invoke the API endpoint
	uri := "https://api.intelligent-api.com/v1/text/classify/multiple"
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", authorization)
	request.Header.Set("User-Agent", userAgent)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Body:", string(body))
}
