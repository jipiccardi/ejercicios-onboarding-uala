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
	return lambdacontext.NewContext(
		context.TODO(),
		&lambdacontext.LambdaContext{
			AwsRequestID:       "awsRequest",
			InvokedFunctionArn: "arn:aws:lambda:xxx",
		},
	)
}
