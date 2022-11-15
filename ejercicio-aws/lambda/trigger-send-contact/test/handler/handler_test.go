package handler

import (
	"context"
	"errors"
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

	happyPathInput := map[string]events.DynamoDBAttributeValue{
		"id":        events.NewStringAttribute("id1234"),
		"firstName": events.NewStringAttribute("first-name"),
		"lastName":  events.NewStringAttribute("last-name"),
	}
	missingIDInput := map[string]events.DynamoDBAttributeValue{
		"id":        events.NewStringAttribute(""),
		"firstName": events.NewStringAttribute("first-name"),
		"lastName":  events.NewStringAttribute("last-name"),
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
			args: args{
				ctx: mocks.Context(),
				e: events.DynamoDBEvent{
					Records: []events.DynamoDBEventRecord{
						{
							Change: events.DynamoDBStreamRecord{
								NewImage: happyPathInput,
							},
						},
					},
				},
			},
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
		{
			name: "error path: process failed",
			args: args{
				ctx: mocks.Context(),
				e: events.DynamoDBEvent{
					Records: []events.DynamoDBEventRecord{
						{
							Change: events.DynamoDBStreamRecord{
								NewImage: happyPathInput,
							},
						},
					},
				},
			},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("Process", dto.Contacto{
					Id:        "id1234",
					FirstName: "first-name",
					LastName:  "last-name",
				}).Return("", errors.New("internal server error"))
			},
			wantErr: assert.Error,
		},
		{
			name: "error path: missing id",
			args: args{
				ctx: mocks.Context(),
				e: events.DynamoDBEvent{
					Records: []events.DynamoDBEventRecord{
						{
							Change: events.DynamoDBStreamRecord{
								NewImage: missingIDInput,
							},
						},
					},
				},
			},
			mock:    mocks.Mock{},
			init:    func(in *mocks.Mock) {},
			wantErr: assert.Error,
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
