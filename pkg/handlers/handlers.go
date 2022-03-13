package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

var ErrorMethodNotAllowed = "method not allowed"

func GetUser() (*events.APIGatewayProxyResponse, error) {

	return nil, nil
}

func CreateUser() (*events.APIGatewayProxyResponse, error) {
	return nil, nil

}

func UpdateUser() (*events.APIGatewayProxyResponse, error) {
	return nil, nil

}

func DeleteUser() (*events.APIGatewayProxyResponse, error) {
	return nil, nil

}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
