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

type Status struct {
	Status string `json:"status"`
}

type Member struct {
	Address string `json:"address"`
	Vars    Vars   `json:"vars"`
}

type Vars struct {
	Token string `json:"token"`
}

func getUser(email string, mailgunBaseURL string, mailgunKey string) (*Member, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/lists/newsletter@tomlinson.email/members/%s", mailgunBaseURL, email), nil)
	req.SetBasicAuth("api", mailgunKey)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if resp.StatusCode >= 300 {
		err = fmt.Errorf("%d: %s", resp.StatusCode, string(body))
		log.Error(err)
		return nil, err
	}
	log.Info(string(body))

	member := Member{}
	json.Unmarshal(body, &member)

	return &member, nil
}

func subscribeUser(email string, mailgunBaseURL string, mailgunKey string) error {

	form := url.Values{}
	form.Add("address", email)
	form.Add("subscribed", "yes")
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

func sendConfirmationEmail(email string, mailgunBaseURL string, mailgunKey string) error {

	form := url.Values{}
	form.Add("from", "Jacob Tomlinson (Newsletter) <jacob+newsletter@tomlinson.email>")
	form.Add("to", email)
	form.Add("subject", "Newsletter: Subscription confirmed")
	form.Add("html", "Thank you for confirming your email address, keep your eyes peeled for your first issue.")

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", mailgunBaseURL, "/tomlinson.email/messages"), strings.NewReader(form.Encode()))
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

	email := request.QueryStringParameters["email"]
	token := request.QueryStringParameters["token"]
	mailgunKey := os.Getenv("MAILGUN_API_KEY")
	mailgunBaseURL := os.Getenv("MAILGUN_BASE_URL")

	log.Info(fmt.Sprintf("Subscribing %s", email))

	member, err := getUser(email, mailgunBaseURL, mailgunKey)
	if err != nil {
		return buildResponse("Unable to get user", 400), nil
	}

	if member.Vars.Token != token {
		log.Info("%s != %s", member.Vars.Token, token)
		return buildResponse("Token does not match", 400), nil
	}

	err = subscribeUser(email, mailgunBaseURL, mailgunKey)
	if err != nil {
		return buildResponse("Unable to subscribe user", 400), nil
	}

	err = sendConfirmationEmail(email, mailgunBaseURL, mailgunKey)
	if err != nil {
		return buildResponse("Unable to send confirmation email", 400), nil
	}

	return buildResponse("Confirmed", 200), nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
