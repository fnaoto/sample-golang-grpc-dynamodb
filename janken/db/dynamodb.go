package db

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateTable(tableName string) {
	dynamodbEndpoint := "http://dynamodb:8000"
	dynamodbRegion := "ap-northeast-1"

	awsSession, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(dynamodbEndpoint),
		Region:   aws.String(dynamodbRegion),
	})

	if err != nil {
		log.Fatal(err)
	}

	ddb := dynamodb.New(awsSession, aws.NewConfig())

	describeTable, err := ddb.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	if describeTable.Table != nil || err != nil {
		fmt.Println(describeTable.Table)
		return
	}

	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(tableName),
	}

	resp, err := ddb.CreateTable(params)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
