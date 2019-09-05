package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var titleRegexp = regexp.MustCompile(`.*`)
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

type news struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Print("CP-MAIN-25")
	switch req.HTTPMethod {
	case "GET":
		return show(req)
	case "POST":
		return create(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func show(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Print("CP-MAIN-37")
	title := req.QueryStringParameters["title"]
	fmt.Print("CP-MAIN-39-TITLE: " + title)
	if !titleRegexp.MatchString(title) {
		return clientError(http.StatusBadRequest)
	}

	n, err := getItem(title)
	if err != nil {
		return serverError(err)
	}
	if n == nil {
		return clientError(http.StatusNotFound)
	}

	js, err := json.Marshal(n)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func create(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Print("CP-MAIN-63")
	if req.Headers["Content-Type"] != "application/json" {
		return clientError(http.StatusNotAcceptable)
	}

	n := new(news)
	err := json.Unmarshal([]byte(req.Body), n)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	if !titleRegexp.MatchString(n.Title) {
		fmt.Print("CP-MAIN-75")
		return clientError(http.StatusBadRequest)
	}
	if n.Description == "" || n.Date == "" {
		fmt.Print("CP-MAIN-79")
		return clientError(http.StatusBadRequest)
	}

	err = putItem(n)
	fmt.Print("CP-MAIN-84")
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Location": fmt.Sprintf("/news?title=%s", n.Title)},
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	fmt.Print("CP-MAIN-96")
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	fmt.Print("CP-MAIN-106")
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(router)
}
