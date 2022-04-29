package main

import (
	"bytes"
	"embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"text/template"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

//go:embed og-template.svg
var f embed.FS

type requestData struct {
	Params map[string]string
}

// Load SVG template
func readSVGTemplate() (*template.Template, error) {
	in, _ := f.ReadFile("og-template.svg")
	return template.New("template").Parse(string(in))
}

// Populate SVG template
func populateSVG(data requestData, tmpl *template.Template) (*bytes.Buffer, error) {
	var out bytes.Buffer
	err := tmpl.Execute(&out, data)
	logParams, _ := json.Marshal(data.Params)
	log.Println(fmt.Sprintf("populating template with data: %s", string(logParams)))
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// Load SVG data into rasteriser
func convertToPNG(svgBuffer *bytes.Buffer) (*bytes.Buffer, error) {
	w, h := 1200, 630
	icon, err := oksvg.ReadIconStream(svgBuffer)
	if err != nil {
		return nil, err
	}
	icon.SetTarget(0, 0, float64(w), float64(h))
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)

	// Convert to PNG
	var buf bytes.Buffer
	err = png.Encode(&buf, rgba)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	tmpl, err := readSVGTemplate()
	if err != nil {
		return nil, err
	}

	svgBuffer, err := populateSVG(requestData{Params: request.QueryStringParameters}, tmpl)
	if err != nil {
		return nil, err
	}

	pngBuffer, err := convertToPNG(svgBuffer)
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "image/png"},
		MultiValueHeaders: http.Header{"Set-Cookie": {"Hello", "World"}},
		Body:              base64.StdEncoding.EncodeToString(pngBuffer.Bytes()),
		IsBase64Encoded:   true,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
