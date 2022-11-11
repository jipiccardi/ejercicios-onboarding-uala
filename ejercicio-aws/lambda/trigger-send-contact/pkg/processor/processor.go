package processor

import (
	"context"
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

func (p *Processor) Process(ctx context.Context, id string) error {

	// TODO: definir como constante
	topic := "arn:aws:sns:us-east-1:620097380428:ContactsTopic-Piccardi"
	res, err := p.SnsClient.PublishMessage(ctx, id, topic)
	if err != nil {
		fmt.Printf("ERR: error publishing the message: %s\n", err)
		return err
	}

	fmt.Printf("Publish message result: %v\n", res)
	return nil

}
