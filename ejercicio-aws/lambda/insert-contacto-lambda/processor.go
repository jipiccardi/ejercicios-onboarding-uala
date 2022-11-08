package main

import (
	"context"

	"github.com/google/uuid"
)

type Contact struct {
	Id        string `dynamodbav:"id"`
	FirstName string `dynamodbav:"firstName"`
	LastName  string `dynamodbav:"lastName"`
	Status    string `dynamodbav:"status"`
}

func Process(ctx context.Context, req InsertContactoRequest) (string, error) {
	item := Contact{
		Id:        uuid.New().String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Status:    "CREATED",
	}

	err := PostContacto(item)
	if err != nil {
		return "", err
	}

	return item.Id, nil
}
