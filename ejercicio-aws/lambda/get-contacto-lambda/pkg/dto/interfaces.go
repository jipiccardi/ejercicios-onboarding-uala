package dto

type Processor interface {
	Process(string) (GetContactoResponse, error)
}

type DynamoClient interface {
	GetContactoById(string, *GetContactoResponse) error
}
