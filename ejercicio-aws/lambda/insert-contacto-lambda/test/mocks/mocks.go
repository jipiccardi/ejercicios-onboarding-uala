package mocks

import (
	"context"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/pkg/dto"
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

func (m *Mock) Process(req dto.InsertContactoRequest) (string, error) {
	args := m.Called(req)
	if err := args.Get(1); err != nil {
		return "", err.(error)
	}
	return args.Get(0).(string), nil
}

func (m *Mock) PostContacto(contacto dto.Contacto) error {
	// Devuelvo sin error ya que si quiero usar m.called tengo el problema de que lo que se le pasa a PostContacto se
	// define adentro de Process y al no coincidir cuando uso in.On tira un error. Investigar
	return nil
}
