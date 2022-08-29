package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Status struct {
	Status string `json:"status"`
}

type Vars struct {
	Token string `json:"token"`
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

func createUser(email string, mailgunBaseURL string, mailgunKey string, token string) error {
	vars, _ := json.Marshal(Vars{token})
	form := url.Values{}
	form.Add("address", email)
	form.Add("vars", string(vars))
	form.Add("subscribed", "no")
	form.Add("upsert", "yes")

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/lists/newsletter@tomlinson.email/members", mailgunBaseURL), strings.NewReader(form.Encode()))
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

func sendVerificationEmail(email string, mailgunBaseURL string, mailgunKey string, token string) error {
	form := url.Values{}
	form.Add("from", "Jacob Tomlinson (Newsletter) <jacob+newsletter@tomlinson.email>")
	form.Add("to", email)
	form.Add("subject", "Newsletter: Verify your email address")
	form.Add("html", fmt.Sprintf("Thank you for subscribing to my newsletter. Before I can add you to the mailing list please click <a href=\"https://jacobtomlinson.dev/.netlify/functions/newsletter-verify?email=%s&token=%s\">here</a> to verify your email address.", url.QueryEscape(email), token))

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
	token := uuid.New().String()

	log.Info(fmt.Sprintf("Subscribing %s", email))

	err := createUser(email, mailgunBaseURL, mailgunKey, token)
	if err != nil {
		return buildResponse("Unable to create user", 400), nil
	}

	err = sendVerificationEmail(email, mailgunBaseURL, mailgunKey, token)
	if err != nil {
		return buildResponse("Unable to send verification email", 400), nil
	}

	log.Info("Subscribed")
	return &events.APIGatewayProxyResponse{
		StatusCode: 302,
		Headers:    map[string]string{"Location": "/newsletter/subscribed/"},
		Body:       "",
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
