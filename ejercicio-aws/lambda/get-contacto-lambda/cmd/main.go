package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/get-contacto-lambda/pkg/external"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/get-contacto-lambda/pkg/handler"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/get-contacto-lambda/pkg/processor"
)

func main() {
	p := processor.New(&external.DynamoClient{})
	lambda.Start(handler.New(p).HandleRequest)
}
