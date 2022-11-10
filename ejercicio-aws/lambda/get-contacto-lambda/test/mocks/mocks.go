package mocks

import (
	"context"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/get-contacto-lambda/pkg/dto"
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

func (m *Mock) Process(id string) (dto.GetContactoResponse, error) {
	args := m.Called(id)
	if err := args.Get(1); err != nil {
		return dto.GetContactoResponse{}, err.(error)
	}
	return args.Get(0).(dto.GetContactoResponse), nil
}

func (m *Mock) GetContactoById(id string, contact *dto.GetContactoResponse) error {
	args := m.Called(id, contact)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
