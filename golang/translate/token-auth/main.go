package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TokenRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func main() {
	// basic configuration
	clientId := "[[client_id]]"
	clientSecret := "[[client_secret]]"
	scopes := "[[scopes]]"

	// call the token endpoint to get an access_token
	tokenRequest := TokenRequest{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
		Scope:        scopes,
	}

	marshalled, err := json.Marshal(tokenRequest)

	if err != nil {
		fmt.Println("Error marshalling token request:", err)
		return
	}

	tokenUri := "https://api.intelligent-api.com/v1/token"
	tokenPostResponse, err := http.Post(tokenUri, "application/json", bytes.NewBuffer([]byte(marshalled)))

	if err != nil {
		fmt.Println("Error marshalling token request:", err)
		return
	}

	if tokenPostResponse.StatusCode != http.StatusOK {
		fmt.Println(fmt.Sprintf("Invalid status code %d returned from %s call", tokenPostResponse.StatusCode, tokenUri))
		return
	}

	responseData, err := io.ReadAll(tokenPostResponse.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(responseData, &tokenResponse)

	if err != nil {
		fmt.Println("Error unmarshalling token response:", err)
		return
	}

	// use the access_token as a bearer token
	authorization := fmt.Sprintf("Bearer %s", tokenResponse.AccessToken)

	// populate endpoint parameters
	textBody := "[[text_body]]"
	sourceLanguage := "[[source_language]]"
	targetLanguage := "[[target_language]]"
	userAgent := "Intelligent API Sample Go Code"

	payload := []byte(fmt.Sprintf(`{ "text": "%s", "sourceLanguage": "%s", "targetLanguage": "%s" }`, textBody, sourceLanguage, targetLanguage))

	// invoke the API endpoint
	uri := "https://api.intelligent-api.com/v1/language/translate"
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
