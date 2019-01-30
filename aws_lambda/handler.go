package main

import (
	"context"
	"github.com/KonstantinAfonin/CalculationsWebService/json"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, calculateRequest json.CalculateRequest) (*json.CalculateResponse, error) {
	return json.Calculate(&calculateRequest)
}

func main() {
	lambda.Start(HandleRequest)
}
