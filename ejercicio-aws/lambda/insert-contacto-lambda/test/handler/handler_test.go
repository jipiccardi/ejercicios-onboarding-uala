package handler

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/pkg/dto"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/pkg/handler"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/test/mocks"
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
						"firstName": "first-name",
						"lastName": "last-name"
					}
					`,
				),
			},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("Process", dto.InsertContactoRequest{
					FirstName: "first-name",
					LastName:  "last-name",
				}).Return("id1234", nil)
			},
			wantValue: assert.NotNil,
			wantErr:   assert.NoError,
		},
		{
			name: "error path: missing field",
			args: args{
				ctx: mocks.Context(),
				payload: json.RawMessage(
					`
					{
						"firstName": "first-name"
					}
					`,
				),
			},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("Process", dto.InsertContactoRequest{}).Return("id1234", nil)
			},
			wantValue: assert.NotNil,
			wantErr:   assert.Error,
		},
		{
			name: "error path: wrong field type",
			args: args{
				ctx: mocks.Context(),
				payload: json.RawMessage(
					`
					{
						"firstName": "first-name",
						"lastName": 324
					}
					`,
				),
			},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("Process", dto.InsertContactoRequest{}).Return("id1234", nil)
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
