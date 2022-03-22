package operations

import (
	"encoding/csv"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"net/http"
	"os"
	"strconv"
	"time"
)

func DownloadResults(apiRoot string, token string) {
	client := &http.Client{}
	// get list of tasks
	tasks, err := getTasks(client, apiRoot, token)
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
	// get submissions in the selected task
	submissions, err := getSubmissionsByTask(client, apiRoot, token, selectedTask.Id, true)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(fmt.Sprintf("%s_results_%s.csv", selectedTask.Name, time.Now().String()))
	defer f.Close()
	w := csv.NewWriter(f)
	err = w.Write([]string{
		"user_id", "username", "point", "notes", "created_at",
	})
	if err != nil {
		panic(err)
	}
	participants, err := getParticipantsByCourse(client, apiRoot, token, selectedTask.CourseId)
	if err != nil {
		panic(err)
	}
	m := make(map[int]string)
	for _, participant := range participants {
		m[participant.User.Id] = participant.User.Username
	}
	for _, submission := range submissions {
		err = w.Write([]string{
			strconv.Itoa(submission.UserId), m[submission.UserId], submission.Point, submission.Notes, submission.CreatedAt,
		})
		if err != nil {
			panic(err)
		}
	}
	w.Flush()
}
