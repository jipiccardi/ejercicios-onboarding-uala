package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-update-contact-lambda/pkg/dto"
)

type Handler struct {
	processor dto.Processor
}

func New(in dto.Processor) *Handler {
	return &Handler{
		processor: in,
	}
}

func (h *Handler) HandleRequest(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		var contacto dto.Contacto
		snsRecord := record.SNS
		if err := json.Unmarshal([]byte(snsRecord.Message), &contacto); err != nil {
			fmt.Printf("ERR: %s\n", err.Error())
		}
		if err := contacto.Validate(); err != nil {
			fmt.Printf("ERR: %s\n", err.Error())
		}

		err := h.processor.Process(contacto.Id)
		if err != nil {
			fmt.Printf("ERR: %s\n", err)
		}

		fmt.Printf("Status updated of contact with id: %s\n", contacto.Id)
	}
}
