package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Status struct {
	Status string `json:"status"`
}

func buildResponse(message string, statusCode int) *events.APIGatewayProxyResponse {
	if statusCode >= 400 {
		log.Error(message)
	} else {
		log.Info(message)
	}
	status, _ := json.Marshal(Status{message})
	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(status),
	}
}

func deleteUser(email string, mailgunBaseURL string, mailgunKey string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/lists/newsletter@tomlinson.email/members/%s", mailgunBaseURL, email), nil)
	req.SetBasicAuth("api", mailgunKey)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Error(err)
		return err
	}

	if resp.StatusCode >= 300 {
		err = fmt.Errorf("%d: %s", resp.StatusCode, string(body))
		log.Error(err)
		return err
	}

	log.Info(string(body))
	return nil
}

func sendConfirmationEmail(email string, mailgunBaseURL string, mailgunKey string) error {
	form := url.Values{}
	form.Add("from", "Jacob Tomlinson (Newsletter) <jacob+newsletter@tomlinson.email>")
	form.Add("to", email)
	form.Add("subject", "Newsletter: Unsubscribed")
	form.Add("html", "You have been unsubscribed, sorry to see you go!")

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/tomlinson.email/messages", mailgunBaseURL), strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("api", mailgunKey)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Error(err)
		return err
	}

	if resp.StatusCode >= 300 {
		err = fmt.Errorf("%d: %s", resp.StatusCode, string(body))
		log.Error(err)
		return err
	}

	log.Info(string(body))
	return nil
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	email, _ := url.QueryUnescape(request.QueryStringParameters["email"])
	mailgunKey := os.Getenv("MAILGUN_API_KEY")
	mailgunBaseURL := os.Getenv("MAILGUN_BASE_URL")

	log.Info(fmt.Sprintf("Unsubscribing %s", email))

	err := deleteUser(email, mailgunBaseURL, mailgunKey)
	if err != nil {
		return buildResponse("Unable to delete user", 400), nil
	}

	err = sendConfirmationEmail(email, mailgunBaseURL, mailgunKey)
	if err != nil {
		return buildResponse("Unable to send confirmation email", 400), nil
	}

	log.Info("Unsubscribed")
	return &events.APIGatewayProxyResponse{
		StatusCode: 302,
		Headers:    map[string]string{"Location": "/newsletter/unsubscribed/"},
		Body:       "",
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
