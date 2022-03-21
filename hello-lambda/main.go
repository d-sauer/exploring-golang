package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

type Response struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func Handler(request Request) (Response, error) {
	fmt.Println("Invoking handler")
	return Response{
		Message:   fmt.Sprintf("Process Request ID: %v", request.ID),
		Timestamp: time.Now().Format(time.RFC3339),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
