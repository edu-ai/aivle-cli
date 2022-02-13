package main

import (
	"aivle-cli/models"
	"aivle-cli/operations"
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	OperationPrintToken          = "print token"
	OperationDownloadSubmissions = "download submissions"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name:     "username",
		Prompt:   &survey.Input{Message: "What is your aiVLE username?"},
		Validate: survey.Required,
	},
	{
		Name:     "password",
		Prompt:   &survey.Password{Message: "Password:"},
		Validate: survey.Required,
	},
	{
		Name: "operation",
		Prompt: &survey.Select{
			Message: "Choose an operation:",
			Options: []string{OperationPrintToken, OperationDownloadSubmissions},
			Default: "print token",
		},
	},
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiRoot := os.Getenv("API_ROOT")
	// the answers will be written to this struct
	answers := struct {
		Username  string `survey:"username"`
		Password  string `survey:"password"`
		Operation string `survey:"operation"`
	}{}

	// perform the questions
	err = survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// try getting the token
	resp, err := http.PostForm(apiRoot+"/dj-rest-auth/login/", url.Values{
		"username": {answers.Username},
		"password": {answers.Password},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Unable to login with status code %d\n", resp.StatusCode)
		return
	}

	tokenResponse := models.TokenResponse{}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// handle operations
	if answers.Operation == OperationPrintToken {
		fmt.Println(tokenResponse.Token)
	} else if answers.Operation == OperationDownloadSubmissions {
		operations.DownloadSubmissions(apiRoot, tokenResponse.Token)
	}
}
