package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

// TODO: Separar en paquetes la estructura del codigo.

func main() {
	lambda.Start(HandleRequest)
}
