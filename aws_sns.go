package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SNSPublish(arn string, region string, subject string, message string) error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}

	svc := sns.New(sess, aws.NewConfig().WithRegion(region))

	params := &sns.PublishInput{
		Message:   aws.String(message),
		Subject:   aws.String(subject),
		TargetArn: aws.String(arn),
	}

	_, err = svc.Publish(params)

	if err != nil {
		return err
	}

	return nil
}
