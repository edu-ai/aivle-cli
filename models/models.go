package models

type TokenResponse struct {
	UserId int    `json:"user"`
	Token  string `json:"key"`
}

type TaskListResponse struct {
	Count    int    `json:"count"`
	Next     *int   `json:"next"`
	Previous *int   `json:"previous"`
	Results  []Task `json:"results"`
}

type Task struct {
	Id       int    `json:"id"`
	CourseId int    `json:"course"`
	Name     string `json:"name"`
}
