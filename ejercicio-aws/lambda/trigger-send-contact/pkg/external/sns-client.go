package external

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/dto"
)

type SnsClient struct {
	Api dto.SNSPublishAPI
}

func (s *SnsClient) PublishMessage(ctx context.Context, message string, topic string) (*sns.PublishOutput, error) {

	input := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(topic),
	}

	result, err := s.Api.Publish(ctx, input)
	if err != nil {
		return result, err
	}

	return result, nil
}
