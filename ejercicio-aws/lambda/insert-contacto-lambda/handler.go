package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
)

type InsertContactoRequest struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Status    string `json:"status"`
}

func HandleRequest(ctx context.Context, payload json.RawMessage) error {
	// Validaciones (validar campos que tienen que venir. El id ya lo valida DynamoDB)
	req := InsertContactoRequest{}
	if err := json.Unmarshal(payload, &req); err != nil {
		log.Printf("ERR: %s\n", err.Error())
		return err
	}

	if err := req.Validate(); err != nil {
		log.Printf("ERR: %s\n", err.Error())
		return err
	}

	// Pegar en la base (processor)
	err := Process(ctx, req)
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return err
	}

	log.Printf("Contact succesfully inserted")
	return nil
}

func (r *InsertContactoRequest) Validate() error {
	if r.FirstName == "" {
		return errors.New("wrong payload: missing firstName field")
	}

	if r.LastName == "" {
		return errors.New("wrong payload: missing lastName field")
	}
	return nil
}
