package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type InsertContactoRequest struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Status    string `json:"status"`
}

type InsertContactoResponse struct {
	Message string `json:"message"`
}

type ErrorMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, payload json.RawMessage) (InsertContactoResponse, error) {

	req := InsertContactoRequest{}

	if err := json.Unmarshal(payload, &req); err != nil {
		log.Printf("ERR: %s\n", err.Error())
		return InsertContactoResponse{}, err
	}

	if err := req.Validate(); err != nil {
		log.Printf("ERR: %s\n", err.Error())
		return InsertContactoResponse{}, err
	}

	id, err := Process(ctx, req)
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return InsertContactoResponse{}, err
	}

	log.Printf("Contact succesfully inserted with id: %s", id)

	return InsertContactoResponse{Message: fmt.Sprintf("Contact succesfully inserted with id: %s", id)}, nil
}

func (r *InsertContactoRequest) Validate() error {
	var errMsg ErrorMessage
	errMsg.Status = 400

	if r.FirstName == "" {
		errMsg.Message = "wrong payload: missing firstName field"
	}

	if r.LastName == "" {
		errMsg.Message = "wrong payload: missing lastName field"
	}

	if len(errMsg.Message) != 0 {
		byteErrMsg, _ := json.Marshal(errMsg)
		return errors.New(string(byteErrMsg))
	}

	return nil
}
