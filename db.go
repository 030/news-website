package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("eu-north-1"))

func getItem(title string) (*news, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("news"),
		Key: map[string]*dynamodb.AttributeValue{
			"title": {
				S: aws.String(title),
			},
		},
	}

	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	n := new(news)
	err = dynamodbattribute.UnmarshalMap(result.Item, n)
	if err != nil {
		return nil, err
	}

	return n, nil
}

func putItem(n *news) error {
	fmt.Print("CP-DB-40")
	input := &dynamodb.PutItemInput{
		TableName: aws.String("news"),
		Item: map[string]*dynamodb.AttributeValue{
			"title": {
				S: aws.String(n.Title),
			},
			"description": {
				S: aws.String(n.Description),
			},
			"date": {
				S: aws.String(n.Date),
			},
		},
	}

	_, err := db.PutItem(input)
	return err
}
