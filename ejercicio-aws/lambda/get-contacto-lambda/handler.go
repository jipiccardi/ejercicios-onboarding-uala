package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
)

// TODO: preguntar si hay otra forma mejor de leer desde go los path parameters

type GetContactoRequest struct {
	Id string `json:"id"`
}

type GetContactoResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Status    string `json:"status"`
}

type ErrorMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, payload json.RawMessage) (GetContactoResponse, error) {
	req := GetContactoRequest{}

	if err := json.Unmarshal(payload, &req); err != nil {
		log.Printf("ERR: %s\n", err.Error())
		return GetContactoResponse{}, err
	}

	if len(req.Id) == 0 {
		return GetContactoResponse{}, ErrorResponse(400, "missing id path parameter /contacts/{id}")
	}

	res, err := Process(ctx, req.Id)
	if err != nil {
		return GetContactoResponse{}, ErrorResponse(500, err.Error())
	}

	return res, nil
}

func ErrorResponse(status int, msg string) error {
	errMsg := ErrorMessage{
		Status:  status,
		Message: msg,
	}

	byteErrMsg, _ := json.Marshal(errMsg)
	return errors.New(string(byteErrMsg))
}
