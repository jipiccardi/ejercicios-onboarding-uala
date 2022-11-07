package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

// TODO: separar en paquetes
// TODO: separar parte de dynamo en un archivo que se llame dynamoClient.go
// TODO: agregar test unitarios

func main() {
	lambda.Start(HandleRequest)
}
