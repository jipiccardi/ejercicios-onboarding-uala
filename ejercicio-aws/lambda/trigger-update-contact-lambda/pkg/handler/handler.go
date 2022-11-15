package handler

import (
	"context"
	"encoding/json"
	"errors"
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

func (h *Handler) HandleRequest(ctx context.Context, snsEvent events.SNSEvent) error {
	errorExist := false

	for _, record := range snsEvent.Records {
		var contacto dto.Contacto
		snsRecord := record.SNS

		if err := json.Unmarshal([]byte(snsRecord.Message), &contacto); err != nil {
			fmt.Printf("ERR: %s\n", err.Error())
			errorExist = true
			continue
		}
		if err := contacto.Validate(); err != nil {
			fmt.Printf("ERR: %s\n", err.Error())
			errorExist = true
			continue
		}

		err := h.processor.Process(contacto.Id)
		if err != nil {
			fmt.Printf("ERR: %s\n", err)
			errorExist = true
		}

	}

	if errorExist {
		return errors.New("error processing some record")

	}
	return nil
}
