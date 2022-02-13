package operations

import (
	"aivle-cli/models"
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"net/http"
)

func DownloadSubmissions(apiRoot string, token string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiRoot+"/api/v1/tasks/", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	req.Header.Set("Authorization", "Token "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	taskList := models.TaskListResponse{}
	err = json.NewDecoder(resp.Body).Decode(&taskList)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var taskNames []string
	for _, task := range taskList.Results {
		taskNames = append(taskNames, task.Name)
	}
	questions := []*survey.Question{
		{
			Name:   "select task",
			Prompt: &survey.Select{Message: "Please select a task:", Options: taskNames},
		},
	}
	answers := struct {
		TaskName int `survey:"select task"`
	}{}
	err = survey.Ask(questions, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Selected task number %d\n", answers.TaskName)
}
