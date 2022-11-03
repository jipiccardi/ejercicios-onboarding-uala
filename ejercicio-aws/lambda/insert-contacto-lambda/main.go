package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

// TODO: Preguntar si es necesario separar en paquetes para no tener en main todo (paquete handler, paquete processor)

func main() {
	lambda.Start(HandleRequest)
}
