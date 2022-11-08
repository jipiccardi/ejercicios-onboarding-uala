package main

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload json.RawMessage
	}

	tests := []struct {
		name     string
		args     args
		initMock func(in *mocks.Mock)
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				ctx: mocks.Context(),
				payload: json.RawMessage(
					`
					{
						"firstName": "firstname",
						"lastName" : "3245"
					}
					`,
				),
			},
			initMock: func(in *mocks.Mock) {
				// ACA HAY Q LLAMAR A LA FUNCION MOCK
				// in.On("Process", mock(), InsertContactoRequest{}).Return("", nil)
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Printf("TEST NAME: %s", tt.name)
			_, gotErr := HandleRequest(tt.args.ctx, tt.args.payload)
			tt.wantErr(t, gotErr)
		})
	}

}
