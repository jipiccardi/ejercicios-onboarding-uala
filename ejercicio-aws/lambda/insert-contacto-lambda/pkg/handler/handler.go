package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/pkg/dto"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/pkg/processor"
)

func HandleRequest(ctx context.Context, payload json.RawMessage) (dto.InsertContactoResponse, error) {

	req := dto.InsertContactoRequest{}

	if err := json.Unmarshal(payload, &req); err != nil {
		fmt.Printf("ERR: %s\n", err.Error())
		return dto.InsertContactoResponse{}, err
	}

	if err := req.Validate(); err != nil {
		fmt.Printf("ERR: %s\n", err.Error())
		return dto.InsertContactoResponse{}, err
	}

	id, err := processor.Process(req)
	if err != nil {
		fmt.Printf("ERR: %s\n", err)
		return dto.InsertContactoResponse{}, err
	}

	fmt.Printf("Contact succesfully inserted with id: %s", id)

	return dto.InsertContactoResponse{Message: fmt.Sprintf("Contact succesfully inserted with id: %s", id)}, nil
}
