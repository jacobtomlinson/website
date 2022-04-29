package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
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

func TestTemplate(t *testing.T) {
	assert := assert.New(t)

	tmpl, err := readSVGTemplate()
	assert.Nil(err)

	params := make(map[string]string)
	params["title"] = "Foo"
	svgBuffer, err := populateSVG(requestData{Params: params}, tmpl)
	assert.Nil(err)
	assert.Contains(svgBuffer.String(), "Foo", "svg should contain parameter title")

}
