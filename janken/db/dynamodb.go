package db

import (
	"fmt"
	"janken/pb"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBConfig struct {
	Client *dynamodb.DynamoDB
	Table  string
}

var ddbc = &DynamoDBConfig{}

func Init() {
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

	CreateTable()
}

func PutItem(item *pb.JankenResult) {
	av, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		log.Fatalf("dynamodbattribute: %s", err)
	}

	log.Print(av)

	_, err = ddbc.Client.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(ddbc.Table),
		Item:      av,
	})

	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
}

func ListTables() []*string {
	resp, err := ddbc.Client.ListTables(&dynamodb.ListTablesInput{})

	if err != nil {
		log.Fatal(err)
	}

	return resp.TableNames
}

func DescribeTable() (*dynamodb.DescribeTableOutput, error) {
	resp, err := ddbc.Client.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: &ddbc.Table,
	})

	return resp, err
}

func CreateTable() {
	describeTable, err := DescribeTable()

	if describeTable != nil && err == nil {
		log.Printf("Already exists: %s", describeTable.Table)
		return
	}

	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("N"),
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
