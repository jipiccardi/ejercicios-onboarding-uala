package processor

import (
	"testing"

	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/pkg/dto"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/pkg/processor"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/insert-contacto-lambda/test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_Process(t *testing.T) {
	type args struct {
		req dto.InsertContactoRequest
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
				req: dto.InsertContactoRequest{
					FirstName: "first-name",
					LastName:  "last-name",
				},
			},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("PostContacto", dto.Contacto{}).Return(nil)
			},
			wantErr:   assert.NoError,
			wantValue: assert.NotNil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				tt.init(&tt.mock)
				p := processor.New(&tt.mock)
				res, err := p.Process(tt.args.req)
				tt.wantErr(t, err)
				tt.wantValue(t, res)
			},
		)
	}

}
