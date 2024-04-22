package main

import (
	"encoding/json"
	"fmt"

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

	var (
		result string
		status int
	)

	switch action {
	case "create":
		name, ok := req["name"].(string)
		if !ok {
			return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Missing name in request"}, nil
		}
		id, err := user.CreateUser(name)
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error creating user: " + err.Error()}, nil
		}
		result = fmt.Sprintf("User created with ID %d", id)
		status = 201 // Created
	case "get":
		id, ok := req["id"].(float64)
		if !ok {
			return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Missing ID in request"}, nil
		}
		u, err := user.GetUser(uint(id))
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error getting user: " + err.Error()}, nil
		}
		result = u.Name
		status = 200 // OK
	default:
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid action in request"}, nil
	}

	respBody, _ := json.Marshal(map[string]string{"result": result})
	return events.APIGatewayProxyResponse{StatusCode: status, Body: string(respBody)}, nil
}

func main() {
	lambda.Start(Handler)
}
