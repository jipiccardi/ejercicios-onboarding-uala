package main

import (
	"context"
)

func Process(ctx context.Context, id string) (GetContactoResponse, error) {
	contacto := GetContactoResponse{}

	err := GetContactoById(id, &contacto)
	if err != nil {
		return contacto, err
	}

	return contacto, nil
}
