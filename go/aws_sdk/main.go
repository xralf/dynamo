package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const TableName = "StartingPointMain"

type User struct {
	PK         string `dynamodbav:"pk" json:"pk"`
	SK         string `dynamodbav:"sk" json:"sk"`
	Email      string `dynamodbav:"email" json:"email"`
	FirstName  string `dynamodbav:"fistName" json:"firstName"`
	MiddleName string `dynamodbav:"middleName" json:"middleName"`
	LastName   string `dynamodbav:"lastName" json:"lastName"`
}

type Task struct {
	PK   string `dynamodbav:"pk" json:"pk"`
	SK   string `dynamodbav:"sk" json:"sk"`
	Name string `dynamodbav:"name" json:"name"`
	Type string `dynamodbav:"type" json:"type"`
}

func CreateTable(ctx context.Context, client *dynamodb.Client) (err error) {
	const tablePath = "StartingPointDynamo.json"

	data, err := ioutil.ReadFile(tablePath)
	if err != nil {
		fmt.Println("CreateTable::Failed to read file at:", tablePath, err)
		return err
	}

	var input dynamodb.CreateTableInput
	err = json.Unmarshal(data, &input)
	if err != nil {
		fmt.Println("CreateTable::Failed to unmarshal json:", err)
		return err
	}

	output, err := client.CreateTable(ctx, &input)
	if err != nil {
		fmt.Println("CreateTable::Failed to create table with error:", err)
		return err
	}

	fmt.Println("CreateTable::Succeeded with output:", *output)

	return
}

func DeleteTable(ctx context.Context, client *dynamodb.Client, tableName string) (err error) {

	input := dynamodb.DeleteTableInput{TableName: aws.String(tableName)}

	output, err := client.DeleteTable(ctx, &input)
	if err != nil {
		fmt.Println("DeleteTable::Failed to create table with error:", err)
		return err
	}

	fmt.Println("DeleteTable::Succeeded with output:", *output)

	return
}

func main() {
	ctx := context.TODO()

	resolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID: "aws",
			//URL:           "http://localhost:8000",
			URL:           "http://localhost:4566",
			SigningRegion: "localhost",
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("localhost"),
		config.WithEndpointResolver(resolver),
	)

	if err != nil {
		fmt.Println("Error loading background ctx failed to load AWS config with error:", err)
		panic(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	//_ = DeleteTable(ctx, client, TableName)

	err = CreateTable(ctx, client)
	if err != nil {
		panic(err)
	}

	out, err := client.ListTables(ctx, &dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println("Failed to list tables with error:", err)
		panic(err)
	}

	for _, k := range out.TableNames {
		fmt.Println("List table key:", k)
	}

	//id, err := uuid.NewUUID()
	//if err != nil {
	//	fmt.Println("Failed to get uuid with error:", err)
	//	panic(err)
	//}
	//
	//i := Task{
	//	PK:   "task#" + id.String(),
	//	SK:   "task#" + id.String(),
	//	Name: "Docker",
	//	Type: "Somewhere out there",
	//}
	//
	//mi, err := attributevalue.MarshalMap(i)
	//if err != nil {
	//	fmt.Println("Failed to marshal item:", i, err)
	//	panic(err)
	//}
	//
	//input := dynamodb.PutItemInput{
	//	Item: mi,
	//	TableName: aws.String(TableName),
	//}
	//
	//piOut, err := client.PutItem(ctx, &input)
	//if err != nil {
	//	fmt.Println("failed to put item:", input, "with error:", err)
	//	panic(err)
	//}
	//
	//fmt.Println("put item successfully:", *piOut)
}
