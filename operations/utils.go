package operations

import (
	"aivle-cli/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func apiGetRequest(apiRoot string, token string, api string) (*http.Request, error) {
	req, err := http.NewRequest("GET", apiRoot+api, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Token "+token)
	return req, nil
}

func getTasks(client *http.Client, apiRoot string, token string) (tasks []models.Task, err error) {
	req, err := apiGetRequest(apiRoot, token, "/api/v1/tasks/")
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	return
}

func getSubmissionsByTask(client *http.Client, apiRoot string, token string, taskId int, markedForGrading bool) (submissions []models.Submission, err error) {
	req, err := apiGetRequest(apiRoot, token, "/api/v1/submissions/")
	if err != nil {
		return
	}
	q := req.URL.Query()
	q.Add("task", strconv.Itoa(taskId))
	if markedForGrading {
		q.Add("marked_for_grading", "true")
	} else {
		q.Add("marked_for_grading", "false")
	}
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&submissions)
	return
}
