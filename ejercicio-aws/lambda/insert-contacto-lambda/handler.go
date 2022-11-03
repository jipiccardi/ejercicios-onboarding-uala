package main

import (
	"context"
	"fmt"
)

func HandleRequest(ctx context.Context) (string, error) {
	// Validaciones (validar campos que tienen que venir. El id ya lo valida DynamoDB)
	// Pegar en la base (processor)

	fmt.Println("Hello from handler!")
	return "Hello from handler!", nil
}
