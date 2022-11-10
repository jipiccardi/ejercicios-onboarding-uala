package handler

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/get-contacto-lambda/pkg/dto"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/get-contacto-lambda/pkg/handler"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/get-contacto-lambda/test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_HandleRequest(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload json.RawMessage
	}

	tests := []struct {
		name      string
		args      args
		mock      mocks.Mock
		init      func(in *mocks.Mock)
		wantValue assert.ValueAssertionFunc
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "happy path",
			args: args{
				ctx: mocks.Context(),
				payload: json.RawMessage(
					`
					{
						"id": "id1234"
					}
					`,
				),
			},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("Process", "id1234").Return(dto.GetContactoResponse{
					Id:        "id1234",
					FirstName: "first-name",
					LastName:  "last-name",
					Status:    "CREATED",
				}, nil)
			},
			wantValue: assert.NotNil,
			wantErr:   assert.NoError,
		},
		{
			name: "error path: invalid id",
			args: args{
				ctx: mocks.Context(),
				payload: json.RawMessage(
					`
					{
						"id": 32553
					}
					`,
				),
			},
			mock:      mocks.Mock{},
			init:      func(in *mocks.Mock) {},
			wantValue: assert.NotNil,
			wantErr:   assert.Error,
		},
		{
			name: "error path: missing id",
			args: args{
				ctx: mocks.Context(),
				payload: json.RawMessage(
					`
					{
						"id": ""
					}
					`,
				),
			},
			mock:      mocks.Mock{},
			init:      func(in *mocks.Mock) {},
			wantValue: assert.NotNil,
			wantErr:   assert.Error,
		},
		{
			name: "error path: internal server error: process failed",
			args: args{
				ctx: mocks.Context(),
				payload: json.RawMessage(
					`
					{
						"id": "id1234"
					}
					`,
				),
			},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("Process", "id1234").Return(dto.GetContactoResponse{}, errors.New("internal server error"))
			},
			wantValue: assert.NotNil,
			wantErr:   assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				tt.init(&tt.mock)
				h := handler.New(&tt.mock)
				res, err := h.HandleRequest(tt.args.ctx, tt.args.payload)
				tt.wantErr(t, err)
				tt.wantValue(t, res)
			},
		)
	}

}
