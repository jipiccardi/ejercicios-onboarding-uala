package main

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetContactoById(id string, contact *GetContactoResponse) error {
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
