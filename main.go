package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

// Full endpoint URL
const openNotifyEndpoint = "http://api.open-notify.org/iss-now.json"

// ISS Position
type issPosition struct {
	latitude  float64 `json:"latitude"`
	longitude float64 `json:"longitude"`
}

// OpenNotify Response
// More at open-notify.org/Open-Notify-API/ISS-Location-Now/
type openNotifyResponse struct {
	message     string      `json:"message"`
	timestamp   int         `json:"timestamp"`
	issPosition issPosition `json:"iss_position"`
}

func main() {

	// Build the request to OpenNotify
	req, err := http.NewRequest("GET", openNotifyEndpoint, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// Http Client
	client := &http.Client{}

	// Making the request
	httpResponse, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	// Closing Body when done reading
	defer httpResponse.Body.Close()

	// Fill the response from OpenNotify
	var response openNotifyResponse

	// Decoding the response
	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		log.Println(err)
	}

	fmt.Println(response)
}
