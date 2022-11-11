package dto

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type Processor interface {
	Process(context.Context, string) error
}

type SnsClient interface {
	PublishMessage(context.Context, string, string) (*sns.PublishOutput, error)
}

type SNSPublishAPI interface {
	Publish(ctx context.Context,
		params *sns.PublishInput,
		optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}
