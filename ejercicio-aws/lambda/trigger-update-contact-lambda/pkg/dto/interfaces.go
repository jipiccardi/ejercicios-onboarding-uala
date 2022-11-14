package dto

type Processor interface {
	Process(id string) error
}

type DynamoClient interface {
	UpdateStatus(string, string) error
}
