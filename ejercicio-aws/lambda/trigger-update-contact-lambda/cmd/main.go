package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-update-contact-lambda/pkg/external"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-update-contact-lambda/pkg/handler"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-update-contact-lambda/pkg/processor"
)

func main() {
	dynamoClient := external.DynamoClient{}
	p := processor.New(&dynamoClient)
	lambda.Start(handler.New(p).HandleRequest)
}
