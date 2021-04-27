package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"image"
	"image/png"
	"os"
	"bytes"
  
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	w, h := 1200, 630

	in, err := os.Open("../../static/og-template.svg")
	if err != nil {
	  panic(err)
	}
	defer in.Close()
  
	icon, _ := oksvg.ReadIconStream(in)
	icon.SetTarget(0, 0, float64(w), float64(h))
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)
  
	var buf bytes.Buffer
  
	err = png.Encode(&buf, rgba)
	if err != nil {
	  panic(err)
	}
	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "image/png"},
		MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
		Body:              buf.String(),
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
