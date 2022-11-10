package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/handler"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/processor"
)

func main() {
	p := processor.New()
	lambda.Start(handler.New(p).HandleRequest)
}
