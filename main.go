package main

import (
	"log"
	"fmt"
	"github.com/pedrolopesme/iss-notifier/iss"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
)

func track() (iss.IssPosition, error) {
	// Retrieving coordinate
	issCoordinate, err := iss.GetCoordinate()
	if err != nil {
		log.Fatal("iss.GetCoordinate: ", err)
	}
	//return fmt.Sprintf("%#v", issCoordinate), err
	return issCoordinate, err
}

func main() {

	issCoordinate, err := track()
	if err  != nil {
		return
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	// URL to our queue
	qURL := "https://sqs.us-east-1.amazonaws.com/183162777455/iss-positions"

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Latitude": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(issCoordinate.Latitude),
			},
			"Longitude": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(issCoordinate.Longitude),
			},
		},
		MessageBody: aws.String("Iss Coordinate"),
		QueueUrl:    &qURL,
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success", *result.MessageId)

	//lambda.Start(track)
}
