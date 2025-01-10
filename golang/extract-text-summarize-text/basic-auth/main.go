package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
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
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Get the file's content type
	buffer := make([]byte, 512) // Read the first 512 bytes
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading file:", err)
		return
	}

	// Detect the content type
	contentType := http.DetectContentType(buffer)

	// Reset file read pointer to the beginning
	file.Seek(0, io.SeekStart)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// invoke the API endpoint
	uri := "https://api.intelligent-api.com/v1/document/text/summarize"
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(fileBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Set("Content-Type", contentType)
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
