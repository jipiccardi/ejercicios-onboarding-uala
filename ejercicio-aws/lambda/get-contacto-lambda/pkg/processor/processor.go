package processor

import (
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/get-contacto-lambda/pkg/dto"
)

type Processor struct {
	DynamoClient dto.DynamoClient
}

func New(client dto.DynamoClient) *Processor {
	return &Processor{
		DynamoClient: client,
	}
}

func (p *Processor) Process(id string) (dto.GetContactoResponse, error) {
	contacto := dto.GetContactoResponse{}

	err := p.DynamoClient.GetContactoById(id, &contacto)
	if err != nil {
		return contacto, err
	}

	return contacto, nil
}
