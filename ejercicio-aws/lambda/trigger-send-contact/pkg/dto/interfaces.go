package dto

import (
	"github.com/aws/aws-sdk-go/service/sns"
)

type Processor interface {
	Process(Contacto) error
}

type SnsClient interface {
	PublishMessage(Contacto) (*sns.PublishOutput, error)
}
