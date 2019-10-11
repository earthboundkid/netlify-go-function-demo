package main

import (
    "fmt"
    "time"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    now := time.Now()
    return &events.APIGatewayProxyResponse{
        StatusCode: 200,
        Body:       fmt.Sprintf("Hello, World %s", now.Format("2006-01-02 3:05:06pm")),
        Headers: map[string]string{
            "Cache-Control": "public, max-age=15",
        },
    }, nil
}

func main() {
    // Make the handler available for Remote Procedure Call by AWS Lambda
    lambda.Start(handler)
}
