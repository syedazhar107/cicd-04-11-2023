package main

import (
	"cicd-03112023-01/internal/mathops"
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	res := mathops.GetSum(10, 20)
	log.Printf("addition is %v", res)
	return "Hello, World!", nil
}

func main() {
	lambda.Start(HandleRequest)
}
