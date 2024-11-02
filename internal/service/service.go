package service

import (
	"api/internal/model"
	"api/internal/repository"
)

type TaskService interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task model.Task) error
	UpdateTask(task model.Task) error
	DeleteTask(id int) error
}

type taskService struct {
	repo repository.TaskRepo
}

func NewService(repo repository.TaskRepo) TaskService {
	return &taskService{repo: repo}
}

func (ts *taskService) GetAllTasks() ([]model.Task, error) {
	return ts.repo.GetAllTasks()
}

func (ts *taskService) CreateTask(task model.Task) error {
	return ts.repo.CreateTask(task)
}


func (ts *taskService) UpdateTask(task model.Task) error {
	return ts.repo.UpdateTask(task)
}

func (ts *taskService) DeleteTask(id int) error {
	return ts.repo.DeleteTask(id)
}
