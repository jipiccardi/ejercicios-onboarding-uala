package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/pkg/handler"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
