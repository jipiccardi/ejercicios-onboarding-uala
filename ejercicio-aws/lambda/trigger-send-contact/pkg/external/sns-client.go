package external

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/dto"
)

type SnsClient struct {
	Api      snsiface.SNSAPI
	TopicARN string
}

func InitSnsClient(topicName string) SnsClient {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return SnsClient{
		Api:      sns.New(sess),
		TopicARN: topicName,
	}
}

func (s *SnsClient) PublishMessage(message dto.Contacto) (*sns.PublishOutput, error) {
	var publishResult *sns.PublishOutput

	byteMsg, err := json.Marshal(message)
	if err != nil {
		return publishResult, err
	}

	input := &sns.PublishInput{
		Message:  aws.String(string(byteMsg)),
		TopicArn: aws.String(s.TopicARN),
	}

	publishResult, err = s.Api.Publish(input)
	if err != nil {
		return publishResult, err
	}

	return publishResult, err
}
