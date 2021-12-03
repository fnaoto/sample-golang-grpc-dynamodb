package db

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBConfig struct {
	Client *dynamodb.DynamoDB
	Table  string
}

var ddbc = &DynamoDBConfig{}

func init() {
	dynamodbEndpoint := "http://dynamodb:8000"
	dynamodbRegion := "ap-northeast-1"
	dynamodbTable := "table"

	awsSession, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(dynamodbEndpoint),
		Region:   aws.String(dynamodbRegion),
	})

	if err != nil {
		log.Fatal(err)
	}

	ddbc = &DynamoDBConfig{
		Client: dynamodb.New(awsSession, aws.NewConfig()),
		Table:  dynamodbTable,
	}
}

func ListTables() []*string {
	resp, err := ddbc.Client.ListTables(&dynamodb.ListTablesInput{})

	if err != nil {
		log.Fatal(err)
	}

	return resp.TableNames
}

func DescribeTable() *dynamodb.TableDescription {
	resp, err := ddbc.Client.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: &ddbc.Table,
	})

	if err != nil {
		log.Fatal(err)
	}

	return resp.Table
}

func CreateTable() {
	describeTable := DescribeTable
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
		TableName: aws.String(ddbc.Table),
	}

	resp, err := ddbc.Client.CreateTable(params)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
