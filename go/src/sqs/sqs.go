package sqs

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQS struct {
	sqs   *sqs.SQS
	queue string
	url   string
}

func NewSQS(queueName string) *SQS {
	sqsEndpoint := os.Getenv("SQS_ENDPOINT")
	// region := os.Getenv("SQS_REGION")
	// awsFlag := strings.Contains(sqsEndpoint, "amazonaws.com")
	// if awsFlag {
	// 	svc := sqs.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	// 	url := "https://sqs.us-east-1.amazonaws.com/478989820108/" + queueName
	// 	return &SQS{
	// 		sqs:   svc,
	// 		queue: queueName,
	// 		url:   url,
	// 	}
	// }
	svc := sqs.New(session.New(), &aws.Config{Endpoint: aws.String(sqsEndpoint), Region: aws.String("us-east-1")})
	url := sqsEndpoint + "/queue/" + queueName
	return &SQS{
		sqs:   svc,
		queue: queueName,
		url:   url,
	}

}

func (svc SQS) createSQSQueue() {
	params := &sqs.CreateQueueInput{
		QueueName: aws.String(svc.queue), // Required
	}
	resp, err := svc.sqs.CreateQueue(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
}

func (svc SQS) getSQSQueueDepth() {

	attrib := "ApproximateNumberOfMessages"
	sendParams := &sqs.GetQueueAttributesInput{
		QueueUrl: aws.String(svc.url), // Required
		AttributeNames: []*string{
			&attrib, // Required
		},
	}
	resp2, sendErr := svc.sqs.GetQueueAttributes(sendParams)
	if sendErr != nil {
		fmt.Println("Depth: " + sendErr.Error())
		return
	}
	fmt.Println(resp2)
}

func (svc SQS) SendMessage(message string) {
	params := &sqs.SendMessageInput{
		MessageBody: aws.String(message), // Required
		QueueUrl:    aws.String(svc.url), // Required
	}
	resp, err := svc.sqs.SendMessage(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)

}

func (svc SQS) ReceiveMessage() {
	params := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(svc.url), // Required
	}
	resp, err := svc.sqs.ReceiveMessage(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
}

func (svc SQS) purgeQueue() {
	params := &sqs.PurgeQueueInput{
		QueueUrl: aws.String(svc.url), // Required
	}
	resp, err := svc.sqs.PurgeQueue(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
}
