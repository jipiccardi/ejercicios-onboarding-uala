package dto

type Contacto struct {
	Id        string `dynamodbav:"id"`
	FirstName string `dynamodbav:"firstName"`
	LastName  string `dynamodbav:"lastName"`
	Status    string `dynamodbav:"status"`
}
