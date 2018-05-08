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
type IssPosition struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// OpenNotify Response
// More at open-notify.org/Open-Notify-API/ISS-Location-Now/
type OpenNotifyResponse struct {
	Message     string      `json:"message"`
	Timestamp   int         `json:"timestamp"`
	IssPosition IssPosition `json:"iss_position"`
}

// getISSCoordinate get satellite coordinate relative to Earth
func getISSCoordinate() (position IssPosition, err error) {

	// Build the request to OpenNotify
	req, err := http.NewRequest("GET", openNotifyEndpoint, nil)
	if err != nil {
		return
	}

	// Http Client
	client := &http.Client{}

	// Making the request
	httpResponse, err := client.Do(req)
	if err != nil {
		return
	}

	// Closing Body when done reading
	defer httpResponse.Body.Close()

	// Fill the response from OpenNotify
	var response OpenNotifyResponse

	// Decoding the response
	if err = json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return
	}

	position = response.IssPosition
	return

}


func main() {

	// Retrieving coordinate
	issCoordinate, err := getISSCoordinate()
	if err != nil {
		log.Fatal("getISSCoordinate: ", err)
	}
	fmt.Println(issCoordinate)
}
