package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Status struct {
	Status string `json:"status"`
}

type Member struct {
	Address    string `json:"address"`
	Vars       string `json:"vars"`
	Subscribed string `json:"subscribed"`
	Upsert     string `json:"upsert"`
}

type Vars struct {
	Token string `json:"token"`
}

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	HTML    string `json:"html"`
}

func createUser(email string, mailgunBaseURL string, mailgunKey string, token string) error {

	vars, _ := json.Marshal(Vars{token})
	jsonValue, _ := json.Marshal(Member{
		Address:    email,
		Vars:       string(vars),
		Subscribed: "no",
		Upsert:     "yes",
	})

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", mailgunBaseURL, "/lists/newsletter@tomlinson.email/members"), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("api", mailgunKey)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if resp.StatusCode >= 300 {
		log.Fatal(err)
		return errors.New(string(body))
	}

	log.Info(string(body))
	return nil
}

func sendVerificationEmail(email string, mailgunBaseURL string, mailgunKey string, token string) error {

	jsonValue, _ := json.Marshal(Message{
		From:    "Jacob Tomlinson (Newsletter) <jacob+newsletter@tomlinson.email",
		To:      email,
		Subject: "Verify your email address",
		HTML:    fmt.Sprintf("Thank you for subscribing to my newsletter. Before I can add you to the mailing list please click <a href=\"https://jacobtomlinson.dev/.netlify/functions/newsletter-verify?email=%s&token=%s\">here</a> to verify your email address.", email, token),
	})

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", mailgunBaseURL, "/tomlinson.email/messages"), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("api", mailgunKey)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if resp.StatusCode >= 300 {
		log.Fatal(err)
		return errors.New(string(body))
	}

	log.Info(string(body))
	return nil
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	email := request.QueryStringParameters["email"]
	mailgunKey := os.Getenv("MAILGUN_API_KEY")
	mailgunBaseURL := os.Getenv("MAILGUN_BASE_URL")
	token := uuid.New().String()

	log.Info(fmt.Sprintf("Subscribing %s", email))

	err := createUser(email, mailgunBaseURL, mailgunKey, token)
	if err != nil {
		log.Fatal("Unable to create user")
		status, _ := json.Marshal(Status{"Unable to create user"})
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       string(status),
		}, nil
	}

	err = sendVerificationEmail(email, mailgunBaseURL, mailgunKey, token)
	if err != nil {
		log.Fatal("Unable to send verification email")
		status, _ := json.Marshal(Status{"Unable to send verification email"})
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       string(status),
		}, nil
	}

	log.Info("Subscribed")
	status, _ := json.Marshal(Status{"Subscribed"})
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(status),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
