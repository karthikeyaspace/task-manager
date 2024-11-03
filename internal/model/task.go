package model

type Task struct {
	ID          int    `json:"id"`
	TaskId      string `json:"taskid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Completed   bool   `json:"completed"`
}
