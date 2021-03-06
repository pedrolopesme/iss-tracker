package sqs

import (
	"github.com/pedrolopesme/iss-tracker/iss"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"encoding/json"
)

func Enqueue(queueUrl string, position iss.IssPosition) (messageId string, err error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	jsonMessage, err := json.Marshal(position)
	if err != nil {
		return
	}

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody: aws.String(string(jsonMessage)),
		QueueUrl:    &queueUrl,
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	messageId =  *result.MessageId
	return
}