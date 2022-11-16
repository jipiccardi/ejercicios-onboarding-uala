package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/dto"
)

type Handler struct {
	processor dto.Processor
}

func New(in dto.Processor) *Handler {
	return &Handler{
		processor: in,
	}
}

// TODO: se podria mejorar para que avise en que registro hubo error y en cuales no.
func (h *Handler) HandleRequest(ctx context.Context, e events.DynamoDBEvent) error {
	errorExist := false

	for _, record := range e.Records {
		fmt.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)

		// En caso de que se modifica el usuario, no hace nada.
		if record.EventName == "MODIFY" {
			continue
		}

		id := record.Change.NewImage["id"].String()
		firstName := record.Change.NewImage["firstName"].String()
		lastName := record.Change.NewImage["lastName"].String()

		if len(id) == 0 {
			fmt.Printf("ERR: missing contact id\n")
			errorExist = true
			continue
		}

		err := h.processor.Process(dto.Contacto{
			Id:        id,
			FirstName: firstName,
			LastName:  lastName,
		})

		if err != nil {
			errorExist = true
		}

	}

	if errorExist {
		return errors.New("error processing some record")
	}

	return nil
}
