package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/pedrolopesme/iss-tracker/iss"
	"github.com/pedrolopesme/iss-tracker/sqs"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	// URL to our queue
	SQSEnvVar = "ISS_SQS_URL"
)

var (
	// ErrSQSVarNotFound is returned when a ENV Var with ISS SQS URL wasn't found
	ErrSQSVarNotFound = errors.New("Environment variable with ISS SQS URL not set. Please define a env var " + SQSEnvVar)
)

// Checks ISS Position and log it into an AWS SQS
func Track() (message string, err error) {

	// Retrieving ISS_SQS_URL from the ENV vars
	queueURL := os.Getenv(SQSEnvVar)
	if queueURL == "" {
		err = ErrSQSVarNotFound
		log.Error(err)
		return
	}

	// Retrieving coordinate
	issCoordinate, err := iss.GetCoordinate()
	if err != nil {
		log.Error("It was impossible to collect satellite position", err)
		return
	}

	// Trying to enqueue to SQS
	message, err = sqs.Enqueue(queueURL, issCoordinate)
	if err != nil {
		log.Error("It was impossible to enqueue the message", err)
		return
	}

	log.Info(
		fmt.Sprintf(
			"Great, iss coordinate recorded! Message id %s, Latitude %s, Longitude %s",
			message,
			issCoordinate.Longitude,
			issCoordinate.Latitude))

	return
}

func main() {
	lambda.Start(Track)
}
