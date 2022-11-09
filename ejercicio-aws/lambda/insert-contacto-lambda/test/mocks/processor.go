package mocks

import (
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/pkg/dto"
)

func (m *Mock) Process(req dto.InsertContactoRequest) (string, error) {
	args := m.Called(req)
	if err := args.Get(1); err != nil {
		return "", err.(error)
	}
	return args.Get(0).(string), nil
}
