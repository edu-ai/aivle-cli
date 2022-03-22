package models

type TokenResponse struct {
	UserId int    `json:"user"`
	Token  string `json:"key"`
}

type Task struct {
	Id       int    `json:"id"`
	CourseId int    `json:"course"`
	Name     string `json:"name"`
}

type Submission struct {
	Id               int    `json:"id"`
	UserId           int    `json:"user"`
	TaskId           int    `json:"task"`
	Point            string `json:"point"`
	Notes            string `json:"notes"`
	CreatedAt        string `json:"created_at"`
	MarkedForGrading bool   `json:"marked_for_grading"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type Participation struct {
	User User   `json:"user"`
	Role string `json:"role"`
}
