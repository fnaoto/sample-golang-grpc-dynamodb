package db

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBConfig struct {
	Client *dynamodb.DynamoDB
	Table  string
}

func (ddbr *DynamoDBConfig) ListTables() []*string {
	resp, err := ddbr.Client.ListTables(&dynamodb.ListTablesInput{})

	if err != nil {
		log.Fatal(err)
	}

	return resp.TableNames
}

func (ddbr *DynamoDBConfig) DescribeTable() *dynamodb.TableDescription {
	resp, err := ddbr.Client.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: &ddbr.Table,
	})

	if err != nil {
		log.Fatal(err)
	}

	return resp.Table
}

func (ddbr *DynamoDBConfig) CreateTable() {
	describeTable := ddbr.DescribeTable
	fmt.Println(&describeTable)

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
		TableName: aws.String(ddbr.Table),
	}

	resp, err := ddbr.Client.CreateTable(params)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
