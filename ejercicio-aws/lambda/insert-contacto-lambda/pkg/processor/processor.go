package processor

import (
	"github.com/google/uuid"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/pkg/dto"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/pkg/external"
)

func Process(req dto.InsertContactoRequest) (string, error) {
	item := dto.Contacto{
		Id:        uuid.New().String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Status:    "CREATED",
	}

	err := external.PostContacto(item)
	if err != nil {
		return "", err
	}

	return item.Id, nil
}
