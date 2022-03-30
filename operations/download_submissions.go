package operations

import (
	"encoding/json"
	"github.com/AlecAivazis/survey/v2"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func DownloadSubmissions(apiRoot string, token string) {
	client := &http.Client{}
	// select one of the tasks
	selectedTask, err := selectOneTask(client, apiRoot, token)
	if err != nil {
		panic(err)
	}
	markedForGrading := true
	err = survey.AskOne(&survey.Confirm{Message: "Download marked-for-grading submissions only? (default Yes)", Default: true}, &markedForGrading)
	// get submissions in the selected task
	submissions, err := getSubmissionsByTask(client, apiRoot, token, selectedTask.Id, markedForGrading)
	if err != nil {
		panic(err)
	}
	// download the submissions
	dirName := ""
	err = survey.AskOne(&survey.Input{Message: "Directory name:"}, &dirName)
	if err != nil {
		panic(err)
	}
	err = os.Mkdir(dirName, 0755)
	if err != nil {
		panic(err)
	}
	for _, submission := range submissions {
		currPath := filepath.Join(dirName, strconv.Itoa(submission.Id))
		err = os.Mkdir(currPath, 0755)
		if err != nil {
			panic(err)
		}
		req, err := apiGetRequest(apiRoot, token, "/api/v1/submissions/"+strconv.Itoa(submission.Id)+"/download/")
		if err != nil {
			panic(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		out, err := os.Create(filepath.Join(currPath, "agent.zip"))
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(err)
		}
		evalResult, err := json.MarshalIndent(submission, "", "  ")
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(filepath.Join(currPath, "result.json"), evalResult, 0644)
		if err != nil {
			panic(err)
		}
	}
}
