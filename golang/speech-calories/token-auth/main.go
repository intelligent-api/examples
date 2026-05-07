package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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
	// token configuration
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
		fmt.Printf("Invalid status code %d returned from %s call\n", tokenPostResponse.StatusCode, tokenUri)
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
	fullPathToFile := "[[filepath]]"
	userAgent := "Intelligent API Sample Go Code"

	// Open the file
	file, err := os.Open(fullPathToFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var fileBuffer bytes.Buffer
	writer := multipart.NewWriter(&fileBuffer)

	part, err := writer.CreateFormFile("file", filepath.Base(fullPathToFile))
	if err != nil {
		fmt.Printf("Error creating form file: %v\n", err)
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Printf("Error copying file content: %v\n", err)
		return
	}

	writer.Close()

	// invoke the API endpoint
	uri := "https://api.intelligent-api.com/v1/speech/calories"
	request, err := http.NewRequest("POST", uri, &fileBuffer)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
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
