package dto

import (
	"encoding/json"
	"errors"
)

type InsertContactoRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
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
