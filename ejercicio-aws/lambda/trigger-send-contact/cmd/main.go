package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/external"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/handler"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/processor"
)

func main() {
	sns := external.InitSnsClient("arn:aws:sns:us-east-1:620097380428:ContactsTopic-Piccardi")
	p := processor.New(&sns)
	lambda.Start(handler.New(p).HandleRequest)
}
