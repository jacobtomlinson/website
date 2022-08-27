package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		MultiValueHeaders: http.Header{"Set-Cookie": {"Hello", "World"}},
		Body:              "Hello world",
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
