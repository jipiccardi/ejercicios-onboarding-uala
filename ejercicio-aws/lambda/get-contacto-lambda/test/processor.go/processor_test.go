package processor

import (
	"errors"
	"testing"

	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/get-contacto-lambda/pkg/dto"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/get-contacto-lambda/pkg/processor"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/get-contacto-lambda/test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_HandleRequest(t *testing.T) {
	type args struct {
		id string
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
			args: args{id: "id1234"},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("GetContactoById", "id1234", &dto.GetContactoResponse{}).Return(nil)
			},
			wantValue: assert.NotNil,
			wantErr:   assert.NoError,
		},
		{
			name: "error path: item not found",
			args: args{id: "id1234"},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("GetContactoById", "id1234", &dto.GetContactoResponse{}).Return(errors.New("item not found"))
			},
			wantValue: assert.NotNil,
			wantErr:   assert.Error,
		},
		{
			name: "error path: error in DynamoDB",
			args: args{id: "id1234"},
			mock: mocks.Mock{},
			init: func(in *mocks.Mock) {
				in.On("GetContactoById", "id1234", &dto.GetContactoResponse{}).Return(errors.New("DynamoDB error"))
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
				p := processor.New(&tt.mock)
				res, err := p.Process(tt.args.id)
				tt.wantErr(t, err)
				tt.wantValue(t, res)
			},
		)
	}

}
