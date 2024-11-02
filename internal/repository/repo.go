package repository

import (
	"api/internal/model"
	"database/sql"
)

type TaskRepo interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task model.Task) error
	GetTask(id int) (model.Task, error)
	UpdateTask(task model.Task) error
	DeleteTask(id int) error
}

type taskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) TaskRepo {
	return &taskRepo{db: db}
}

func (repo *taskRepo) GetAllTasks() ([]model.Task, error) {
	rows, err := repo.db.Query("SELECT id, title, description, completed, priority FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (repo *taskRepo) CreateTask(task model.Task) error {
	_, err := repo.db.Exec("INSERT INTO tasks (title, description, completed, priority) VALUES ($1, $2, $3, $4)", task.Title, task.Description, task.Completed, task.Priority)
	if err != nil {
		return err
	}
	return nil
}

func (repo *taskRepo) GetTask(id int) (model.Task, error) {
	row, err := repo.db.Query("SELECT title, description, completed, priority from tasks WHERE id = $1", id)
	if err != nil {
		return model.Task{}, err
	}
	var task model.Task
	row.Scan(&task.Title, &task.Description, &task.Completed, &task.Priority)
	task.ID = id
	return task, nil
}

func (repo *taskRepo) UpdateTask(task model.Task) error {
	_, err := repo.db.Exec("UPDATE tasks SET title = $1, description = $2, completed = $3, priority = $4 WHERE id = $5", task.Title, task.Description, task.Completed, task.Priority, task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *taskRepo) DeleteTask(id int) error {
	_, err := repo.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
