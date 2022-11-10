package handler

import (
	"context"
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

func (h *Handler) HandleRequest(ctx context.Context, e events.DynamoDBEvent) {

	for _, record := range e.Records {
		fmt.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)

		id := record.Change.NewImage["id"].String()
		firstName := record.Change.NewImage["firstName"].String()
		lastName := record.Change.NewImage["lastName"].String()

		fmt.Printf("contact id: %s, firstName: %s, lastName: %s\n", id, firstName, lastName)
	}
}
