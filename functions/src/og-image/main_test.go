package main

import (
	"testing"
    "github.com/stretchr/testify/assert"
	"github.com/aws/aws-lambda-go/events"
)

func TestOGImage(t *testing.T) {
	assert := assert.New(t)

	params := make(map[string]string)
	params["title"] = "Foo"
	req := events.APIGatewayProxyRequest{QueryStringParameters: params}

	response, err := handler(req)
	assert.Nil(err)
	assert.Equal(200, response.StatusCode, "should return status 200")
}
