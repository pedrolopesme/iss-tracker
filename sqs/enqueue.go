package sqs

import (
	"github.com/pedrolopesme/iss-tracker/iss"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
)

func Enqueue(queueUrl string, position iss.IssPosition) (messageId string, err error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Latitude": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(position.Latitude),
			},
			"Longitude": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(position.Longitude),
			},
		},
		MessageBody: aws.String("ISS Coordinate"),
		QueueUrl:    &queueUrl,
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	messageId =  *result.MessageId
	return
}