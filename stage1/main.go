package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
    "net/http"
    "time"
)

// Request represents the input event for the Lambda function.
type Request struct {
    SlackName string `json:"Ezemba Marvellous"`
    Track     string `json:"Backend"`
}

// Response represents the output of the Lambda function.
type Response struct { 
    SlackName     string `json:"slack_name"`
    CurrentDay    string `json:"current_day"`
    UTCTime       string `json:"utc_time"`
    Track         string `json:"track"`
    GithubFileURL string `json:"github_file_url"`
    GithubRepoURL string `json:"github_repo_url"`
    StatusCode    int    `json:"status_code"`
}

// Handler is your Lambda function's entry point.
func Handler(ctx context.Context, request Request) (Response, error) {
    // Validate parameters
    if request.SlackName == "" || request.Track == "" {
        return Response{}, fmt.Errorf("slack_name and track are required")
    }

    // Get current day of the week
    currentDay := time.Now().Weekday().String()

    // Get current UTC time with validation
    currentTime := time.Now().UTC()
    currentTimeFormatted := currentTime.Format("2006-01-02T15:04:05Z")

    // Construct GitHub URLs
    githubFileURL := "https://github.com/marviigrey/HNG-BACKEND-INTERNSHIP/blob/main/stage1/main.go"
    githubRepoURL := "https://github.com/marviigrey/HNG-BACKEND-INTERNSHIP"

    // Construct the response JSON
    response := Response{
        SlackName:     request.SlackName,
        CurrentDay:    currentDay,
        UTCTime:       currentTimeFormatted,
        Track:         request.Track,
        GithubFileURL: githubFileURL,
        GithubRepoURL: githubRepoURL,
        StatusCode:    http.StatusOK,
    }

    return response, nil
}

func main() {
    lambda.Start(Handler)
}
