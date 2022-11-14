package handler

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/dto"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/handler"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_HandleRequest(t *testing.T) {
	type args struct {
		ctx context.Context
		e   events.DynamoDBEvent
	}

	tests := []struct {
		name    string
		args    args
		mock    mocks.Mock
		init    func(in *mocks.Mock)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "happy path",
			args: args{},
			// TODO: ver como en args pasar el events.DynamoDBEvent.
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("Process", dto.Contacto{
					Id:        "id1234",
					FirstName: "first-name",
					LastName:  "last-name",
				}).Return("id1234", nil)
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				tt.init(&tt.mock)
				h := handler.New(&tt.mock)
				err := h.HandleRequest(tt.args.ctx, tt.args.e)
				tt.wantErr(t, err)
			},
		)
	}
}
