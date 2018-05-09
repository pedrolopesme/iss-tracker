package main

import (
	"log"
	"fmt"
	"github.com/pedrolopesme/iss-notifier/iss"
	"github.com/aws/aws-lambda-go/lambda"
)

func track() (string, error) {
	// Retrieving coordinate
	issCoordinate, err := iss.GetCoordinate()
	if err != nil {
		log.Fatal("iss.GetCoordinate: ", err)
	}
	return fmt.Sprintf("%#v", issCoordinate), err
}

func main() {
	lambda.Start(track)
}
