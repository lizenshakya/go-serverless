package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lizenshakya/go-serverless/internal/user"
)

func Handler(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	var req map[string]interface{}
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid request body"}, err
	}

	action, ok := req["action"].(string)
	if !ok {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Missing action in request"}, nil
	}

	switch action {
	case "create":
		name, ok := req["name"].(string)
		if !ok {
			return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Missing name in request"}, nil
		}
		response, err = user.CreateUser(name)
	case "get":
		id, ok := req["id"].(float64)
		if !ok {
			return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Missing ID in request"}, nil
		}
		response, err = user.GetUser(uint(id))
	}
}

func main() {
	lambda.Start(Handler)
}
