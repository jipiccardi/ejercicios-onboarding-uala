package mocks

import (
	"context"

	"github.com/aws/aws-lambda-go/lambdacontext"
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

func (m *Mock) Process(id string) error {
	args := m.Called(id)
	if err := args.Get(0); err != nil {
		return err.(error)
	}
	return nil
}

func (m *Mock) UpdateStatus(id string, status string) error {
	args := m.Called(id, "PROCESSED")
	if err := args.Get(0); err != nil {
		return err.(error)
	}
	return nil
}
