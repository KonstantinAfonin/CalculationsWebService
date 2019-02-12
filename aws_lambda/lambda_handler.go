package main

import (
	"context"
	"github.com/KonstantinAfonin/CalculatorLambdaHandler/json"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func HandleRequest(ctx context.Context, calculateRequest json.CalculateRequest) (*json.CalculateResponse, error) {
	log.Printf("request: %v", calculateRequest)
	response, err := json.Calculate(&calculateRequest)
	log.Printf("response: %v, error: %v", response, err)
	return response, err
}

func main() {
	lambda.Start(HandleRequest)
}
