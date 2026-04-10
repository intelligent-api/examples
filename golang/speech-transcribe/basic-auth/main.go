package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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
	uri := "https://api.intelligent-api.com/v1/speech/transcribe"
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
