package operations

import (
	"aivle-cli/models"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"net/http"
)

func selectOneTask(client *http.Client, apiRoot string, token string) (selectedTask models.Task, err error) {
	tasks, err := getTasks(client, apiRoot, token)
	if err != nil {
		return
	}
	var taskNames []string
	for _, task := range tasks {
		taskNames = append(taskNames, task.Name)
	}
	var taskIndex = 0
	err = survey.AskOne(&survey.Select{Message: "Please select a task:", Options: taskNames}, &taskIndex)
	if err != nil {
		panic(err)
	}
	selectedTask = tasks[taskIndex]
	return
}

func selectOneCourse(client *http.Client, apiRoot string, token string) (selectedCourse models.Course, err error) {
	courses, err := getCourses(client, apiRoot, token)
	if err != nil {
		return
	}
	var courseNames []string
	for _, course := range courses {
		courseNames = append(courseNames, fmt.Sprintf("%s - %s Semester %d", course.Code, course.AcademicYear, course.Semester))
	}
	var courseIndex = 0
	err = survey.AskOne(&survey.Select{Message: "Please select a course:", Options: courseNames}, &courseIndex)
	if err != nil {
		panic(err)
	}
	selectedCourse = courses[courseIndex]
	return
}
