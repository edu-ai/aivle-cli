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

type Course struct {
	Id            int    `json:"id"`
	Code          string `json:"code"`
	AcademicYear  string `json:"academic_year"`
	Semester      int    `json:"semester"`
	Visible       bool   `json:"visible"`
	Participation struct {
		ParticipationId int    `json:"id"`
		Role            string `json:"role"`
		UserId          int    `json:"user"`
		CourseId        int    `json:"course"`
	} `json:"participation"`
}
