package iss

import (
	"net/http"
	"encoding/json"
	"errors"
)

// Errors
var (
	// ErrOpenNotifyDown is an error returned when it was impossible to connect open notify
	ErrOpenNotifyDown = errors.New("open notify is down")
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

// GetCoordinate get satellite coordinate relative to Earth
func GetCoordinate() (position IssPosition, err error) {

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

	// Check open notify status
	if httpResponse.StatusCode != http.StatusOK {
		err = ErrOpenNotifyDown
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