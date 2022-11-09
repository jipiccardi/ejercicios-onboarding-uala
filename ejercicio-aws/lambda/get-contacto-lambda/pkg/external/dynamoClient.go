package external

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jipiccardi/ejercicios-onboarding-uala/ejercicio-aws/lambda/get-contacto-lambda/pkg/dto"
)

type DynamoClient struct {
}

func (c *DynamoClient) GetContactoById(id string, contact *dto.GetContactoResponse) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	getItem := &dynamodb.GetItemInput{
		TableName: aws.String("Piccardi-Contacts"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := svc.GetItem(getItem)
	if err != nil {
		log.Printf("Got error calling GetItem: %v\n", err)
		return err
	}
	if result.Item == nil {
		return errors.New("item not found")
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, contact)
	return err

}
