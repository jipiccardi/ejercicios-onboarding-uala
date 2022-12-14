package mocks

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/dto"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func Context() context.Context {
	ctx := &lambdacontext.LambdaContext{
		AwsRequestID:       "awsRequestId1234",
		InvokedFunctionArn: "arn:aws:lambda:xxx",
		Identity:           lambdacontext.CognitoIdentity{},
		ClientContext:      lambdacontext.ClientContext{},
	}

	return lambdacontext.NewContext(context.TODO(), ctx)
}

func (m *Mock) Process(contacto dto.Contacto) error {
	args := m.Called(contacto)
	if err := args.Get(1); err != nil {
		fmt.Printf("%+v\n", err)
		return err.(error)
	}
	return nil
}

func (m *Mock) PublishMessage(message dto.Contacto) (*sns.PublishOutput, error) {
	args := m.Called(message)
	if args.Get(1) != nil {
		return &sns.PublishOutput{}, args.Error(1)
	}
	return args.Get(0).(*sns.PublishOutput), nil
}
