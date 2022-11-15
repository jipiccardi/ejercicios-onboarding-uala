package processor

import (
	"errors"
	"testing"

	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-update-contact-lambda/pkg/processor"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-update-contact-lambda/test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_HandleRequest(t *testing.T) {
	type args struct {
		id string
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
			args: args{id: "id1234"},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("UpdateStatus", "id1234", "PROCESSED").Return(nil)
			},
			wantErr: assert.NoError,
		},
		{
			name: "error path",
			args: args{id: "id1234"},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("UpdateStatus", "id1234", "PROCESSED").Return(errors.New("internal server error"))
			},
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init(&tt.mock)
			p := processor.New(&tt.mock)
			err := p.Process(tt.args.id)
			tt.wantErr(t, err)
		})
	}
}
