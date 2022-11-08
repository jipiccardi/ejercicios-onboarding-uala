package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func PostContacto(contacto Contact) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	putItem, err := dynamodbattribute.MarshalMap(contacto)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Piccardi-Contacts"),
		Item:      putItem,
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
