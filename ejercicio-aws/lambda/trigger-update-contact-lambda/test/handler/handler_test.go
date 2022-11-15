package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-update-contact-lambda/pkg/handler"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-update-contact-lambda/test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_HandleRequest(t *testing.T) {
	type args struct {
		ctx      context.Context
		snsEvent events.SNSEvent
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
				snsEvent: events.SNSEvent{
					Records: []events.SNSEventRecord{
						{
							SNS: events.SNSEntity{
								Message: `{"id":"id1234"}`,
							},
						},
					},
				},
			},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("Process", "id1234").Return(nil)
			},
			wantErr: assert.NoError,
		},
		{
			name: "error path: missing id",
			args: args{
				ctx: mocks.Context(),
				snsEvent: events.SNSEvent{
					Records: []events.SNSEventRecord{
						{
							SNS: events.SNSEntity{
								Message: `{"id":""}`,
							},
						},
					},
				},
			},
			mock:    mocks.Mock{},
			init:    func(in *mocks.Mock) {},
			wantErr: assert.Error,
		},
		{
			name: "error path",
			args: args{
				ctx: mocks.Context(),
				snsEvent: events.SNSEvent{
					Records: []events.SNSEventRecord{
						{
							SNS: events.SNSEntity{
								Message: `{"id":"id1234"}`,
							},
						},
					},
				},
			},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("Process", "id1234").Return(errors.New("internal server error"))
			},
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init(&tt.mock)
			h := handler.New(&tt.mock)
			err := h.HandleRequest(tt.args.ctx, tt.args.snsEvent)
			tt.wantErr(t, err)
		})
	}
}
