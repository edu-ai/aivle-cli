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
	OperationExit                = "exit"
	OperationPrintToken          = "print token"
	OperationDownloadSubmissions = "download submissions"
	OperationDownloadResults     = "download results"
)

// the questions to ask
var loginQuestions = []*survey.Question{
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
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiRoot := os.Getenv("API_ROOT")
	// the loginAnswers will be written to this struct
	loginAnswers := struct {
		Username string `survey:"username"`
		Password string `survey:"password"`
	}{}

	// perform the questions
	err = survey.Ask(loginQuestions, &loginAnswers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// try getting the token
	resp, err := http.PostForm(apiRoot+"/dj-rest-auth/login/", url.Values{
		"username": {loginAnswers.Username},
		"password": {loginAnswers.Password},
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

	for {
		// handle operations
		operation := ""
		err = survey.AskOne(&survey.Select{
			Message: "Choose an operation:",
			Options: []string{OperationExit, OperationPrintToken, OperationDownloadSubmissions, OperationDownloadResults},
			Default: "print token",
		}, &operation)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if operation == OperationExit {
			break
		} else if operation == OperationPrintToken {
			fmt.Println(tokenResponse.Token)
		} else if operation == OperationDownloadSubmissions {
			operations.DownloadSubmissions(apiRoot, tokenResponse.Token)
		} else if operation == OperationDownloadResults {
			operations.DownloadResults(apiRoot, tokenResponse.Token)
		}
	}
}
