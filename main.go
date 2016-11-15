package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	var values map[string]string
	json.Unmarshal(evt, &values)
	log.Printf("arn: %s, region: %s, subject: %s,  message: %s\n", values["arn"], values["region"], values["subject"], values["message"])

	err := SNSPublish(values["arn"], values["region"], values["subject"], values["message"])
	if err != nil {
		return nil, err
	}
	log.Println("SNS Published")

	return nil, nil
}

func init() {
	runtime.HandleFunc(handle)
}

func SNSPublish(arn string, region string, subject string, message string) error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}

	params := &sns.PublishInput{
		Message:   aws.String(message),
		Subject:   aws.String(subject),
		TargetArn: aws.String(arn),
	}

	svc := sns.New(sess, aws.NewConfig().WithRegion(region))
	_, err = svc.Publish(params)
	if err != nil {
		return err
	}

	return nil
}

func main() {}
