package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		fmt.Println("Error creating AWS session:", err)
		return
	}

	// Create a DynamoDB service client
	svc := dynamodb.New(sess)

	// Call the ListTables operation
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println("Error listing DynamoDB tables:", err)
		return
	}

	fmt.Println("DynamoDB tables:")
	for _, n := range result.TableNames {
		fmt.Println("-", *n)
	}
}
