package repository

import (
	"api/internal/model"
	"database/sql"
)

type TaskRepo interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task model.Task) error
	UpdateTask(task model.Task) error
	DeleteTask(taskid string) error
}

type taskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) TaskRepo {
	return &taskRepo{db: db}
}

func (repo *taskRepo) GetAllTasks() ([]model.Task, error) {
	rows, err := repo.db.Query("SELECT taskid, title, description, completed, priority FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.TaskId, &task.Title, &task.Description, &task.Completed, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (repo *taskRepo) CreateTask(task model.Task) error {
	_, err := repo.db.Exec("INSERT INTO tasks (taskid, title, description, completed, priority) VALUES ($1, $2, $3, $4, $5)", task.TaskId, task.Title, task.Description, task.Completed, task.Priority)
	if err != nil {
		return err
	}
	return nil
}

func (repo *taskRepo) UpdateTask(task model.Task) error {
	_, err := repo.db.Exec("UPDATE tasks SET title = $1, description = $2, completed = $3, priority = $4 WHERE taskid = $5", task.Title, task.Description, task.Completed, task.Priority, task.TaskId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *taskRepo) DeleteTask(taskid string) error {
	_, err := repo.db.Exec("DELETE FROM tasks WHERE taskid = $1", taskid)
	if err != nil {
		return err
	}
	return nil
}
