package main

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type Contact struct {
	Id        string `dynamodbav:"id"`
	FirstName string `dynamodbav:"firstName"`
	LastName  string `dynamodbav:"lastName"`
	Status    string `dynamodbav:"status"`
}

func Process(ctx context.Context, req InsertContactoRequest) (string, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	item := Contact{
		Id:        uuid.New().String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Status:    "CREATED",
	}
	putItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return "", err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Piccardi-Contacts"),
		Item:      putItem,
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return "", err
	}

	return item.Id, nil
}
