package operations

import (
	"aivle-cli/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func apiGetRequest(apiRoot string, token string, api string) (*http.Request, error) {
	req, err := http.NewRequest("GET", apiRoot+api, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Token "+token)
	return req, nil
}

func apiPostRequest(apiRoot string, token string, api string, formData *url.Values) (*http.Request, error) {
	req, err := http.NewRequest("POST", apiRoot+api, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
	defer resp.Body.Close()
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
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&submissions)
	return
}

func getParticipantsByCourse(client *http.Client, apiRoot string, token string, courseId int) (participants []models.Participation, err error) {
	req, err := apiGetRequest(apiRoot, token, fmt.Sprintf("/api/v1/courses/%d/get_participants/", courseId))
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&participants)
	return
}

func addWhitelist(client *http.Client, apiRoot string, token string, courseId int, email string) (err error) {
	req, err := apiPostRequest(apiRoot, token, "/api/v1/whitelist/",
		&url.Values{"course": {strconv.Itoa(courseId)}, "email": {email}})
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	return
}

func getCourses(client *http.Client, apiRoot string, token string) (courses []models.Course, err error) {
	// TODO: support pagination
	req, err := apiGetRequest(apiRoot, token, "/api/v1/courses/")
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	var results struct {
		Results []models.Course `json:"results"`
	}
	err = json.NewDecoder(resp.Body).Decode(&results)
	return results.Results, err
}
