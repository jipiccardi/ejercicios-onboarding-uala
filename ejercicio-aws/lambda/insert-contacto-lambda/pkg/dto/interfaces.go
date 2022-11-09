package dto

type Processor interface {
	Process(req InsertContactoRequest) (string, error)
}

type DynamoClient interface {
	PostContacto(contacto Contacto) error
}
