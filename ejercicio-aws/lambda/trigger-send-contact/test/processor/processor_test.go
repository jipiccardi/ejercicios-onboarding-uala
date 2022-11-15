package processor

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/dto"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/pkg/processor"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/trigger-send-contact/test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_Process(t *testing.T) {
	type args struct {
		contact dto.Contacto
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
				contact: dto.Contacto{
					Id:        "id1234",
					FirstName: "first-name",
					LastName:  "last-name",
				},
			},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("PublishMessage", dto.Contacto{
					Id:        "id1234",
					FirstName: "first-name",
					LastName:  "last-name",
				}).Return(&sns.PublishOutput{}, nil)
			},
			wantErr: assert.NoError,
		},
		{
			name: "error path",
			args: args{
				contact: dto.Contacto{
					Id:        "id1234",
					FirstName: "first-name",
					LastName:  "last-name",
				},
			},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("PublishMessage", dto.Contacto{
					Id:        "id1234",
					FirstName: "first-name",
					LastName:  "last-name",
				}).Return(&sns.PublishOutput{}, errors.New("internal server error"))
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				tt.init(&tt.mock)
				p := processor.New(&tt.mock)
				err := p.Process(tt.args.contact)
				tt.wantErr(t, err)
			},
		)
	}
}
