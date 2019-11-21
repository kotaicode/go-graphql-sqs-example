package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/rs/zerolog/log"
)

var sqsService *sqs.SQS
var sqsUrl string

func initSqsService(queueName string) error {
	var err error
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	sqsService = sqs.New(sess)

	sqsUrl, err = queueUrl(sqsService, queueName)
	if err != nil {
		log.Err(err)
		return err
	}
	log.Info().Msgf("SQS queue init done: %s", sqsUrl)
	return nil
}

func sendSqsMessage(msg string) (string, error) {
	result, err := sqsService.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Message": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(msg),
			},
		},
		MessageBody: aws.String(msg),
		QueueUrl:    &sqsUrl,
	})

	if err != nil {
		log.Err(err)
		return "", err
	}
	log.Info().Str("SQS message id", *result.MessageId)
	return *result.MessageId, nil
}

func receiveSqsMessage() (string, error) {
	result, err := sqsService.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &sqsUrl,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(20), // 20 seconds
		WaitTimeSeconds:     aws.Int64(0),
	})

	if err != nil {
		log.Err(err)
		return "", err
	}

	if len(result.Messages) == 0 {
		log.Info().Msg("Received no messages")
		return "", nil
	}

	if len(result.Messages) > 0 {
		//log.Info().Msg(result.Messages[0].GoString())
		log.Info().Msgf("%s: %s", *result.Messages[0].MessageId, *result.Messages[0].Body)
	}

	resultDelete, err := sqsService.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &sqsUrl,
		ReceiptHandle: result.Messages[0].ReceiptHandle,
	})

	if err != nil {
		log.Err(err)
		return "", err
	}

	log.Info().Msgf("Message Deleted", resultDelete)

	return *result.Messages[0].MessageId + ": " + *result.Messages[0].Body, nil
}

func queueUrl(svc *sqs.SQS, name string) (string, error) {
	resultURL, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == sqs.ErrCodeQueueDoesNotExist {
			log.Err(err).Msgf("Unable to find queue %q.", name)
		}
		log.Err(err).Msgf("Unable to queue %q, %v.", name, err)
		return "", err
	}
	return *resultURL.QueueUrl, nil
}
