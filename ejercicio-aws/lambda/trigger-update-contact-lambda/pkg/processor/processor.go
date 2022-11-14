package processor

import "github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-update-contact-lambda/pkg/dto"

type Processor struct {
	DynamoClient dto.DynamoClient
}

func New(client dto.DynamoClient) *Processor {
	return &Processor{
		DynamoClient: client,
	}
}

func (p *Processor) Process(id string) error {
	err := p.DynamoClient.UpdateStatus(id, "PROCESSED")
	return err
}
