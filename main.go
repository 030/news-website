package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type news struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

func show() (*news, error) {
	n, err := getItem("some-title")
	if err != nil {
		return nil, err
	}

	return n, nil
}

func main() {
	lambda.Start(show)
}
