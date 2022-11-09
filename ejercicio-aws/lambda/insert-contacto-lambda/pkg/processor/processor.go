package processor

import (
	"github.com/google/uuid"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/pkg/dto"
)

type Processor struct {
	DynamoClient dto.DynamoClient
}

func New(client dto.DynamoClient) *Processor {
	return &Processor{
		DynamoClient: client,
	}
}

func (p *Processor) Process(req dto.InsertContactoRequest) (string, error) {
	item := dto.Contacto{
		Id:        uuid.New().String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Status:    "CREATED",
	}

	err := p.DynamoClient.PostContacto(item)
	if err != nil {
		return "", err
	}

	return item.Id, nil
}
