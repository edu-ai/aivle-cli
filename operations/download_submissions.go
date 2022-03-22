package operations

import (
	"aivle-cli/models"
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
	// get list of tasks
	req, err := apiGetRequest(apiRoot, token, "/api/v1/tasks/")
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	var tasks []models.Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		panic(err)
	}
	// select one of the tasks
	var taskNames []string
	for _, task := range tasks {
		taskNames = append(taskNames, task.Name)
	}
	taskName := 0
	err = survey.AskOne(&survey.Select{Message: "Please select a task:", Options: taskNames}, &taskName)
	if err != nil {
		panic(err)
	}
	selectedTask := tasks[taskName]
	markedForGrading := true
	err = survey.AskOne(&survey.Confirm{Message: "Download marked-for-grading submissions only? (default Yes)", Default: true}, &markedForGrading)
	// get submissions in the selected task
	req, err = apiGetRequest(apiRoot, token, "/api/v1/submissions/")
	if err != nil {
		panic(err)
	}
	q := req.URL.Query()
	q.Add("task", strconv.Itoa(selectedTask.Id))
	if markedForGrading {
		q.Add("marked_for_grading", "true")
	} else {
		q.Add("marked_for_grading", "false")
	}
	req.URL.RawQuery = q.Encode()
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	var submissions []models.Submission
	err = json.NewDecoder(resp.Body).Decode(&submissions)
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
		req, err = apiGetRequest(apiRoot, token, "/api/v1/submissions/"+strconv.Itoa(submission.Id)+"/download/")
		if err != nil {
			panic(err)
		}
		resp, err = client.Do(req)
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
