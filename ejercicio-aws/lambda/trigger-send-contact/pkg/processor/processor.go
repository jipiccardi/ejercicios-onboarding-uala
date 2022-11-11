package processor

import (
	"fmt"

	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/dto"
)

type Processor struct {
	SnsClient dto.SnsClient
}

func New(client dto.SnsClient) *Processor {
	return &Processor{
		SnsClient: client,
	}
}

func (p *Processor) Process(contact dto.Contacto) error {

	res, err := p.SnsClient.PublishMessage(contact)
	if err != nil {
		fmt.Printf("ERR: error publishing the message: %s\n", err)
	}

	fmt.Printf("Publish message result: %v\n", res)

	return nil
}
