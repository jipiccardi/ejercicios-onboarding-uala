package mocks

import "context"

// TODO: ordenar el codigo en paquetes es inevitable... es mas facil laburarlo asi.
// TODO: entender y probar el concepto mock

// TODO: insertContactoRequest deberia estar en una carpeta que se llama dto donde estan las estructuras

// CMD -> main.go pkg-> handler,dto,

func (m *Mock) Process(ctx context.Context, req InsertContactoRequest) (dto.Response, error) {
	args := m.Called(ctx, in)
	if err := args.Get(1); err != nil {
		return nil, err.(error)
	}
	return args.Get(0), nil
}
