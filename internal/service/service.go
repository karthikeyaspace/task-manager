package service

import (
	"api/internal/model"
	"api/internal/repository"
	"github.com/google/uuid"
)

type TaskService interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task model.Task) (string, error)
	UpdateTask(task model.Task) error
	DeleteTask(taskid string) error
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

func (ts *taskService) CreateTask(task model.Task) (string, error) {
	tid := uuid.New().String()
	task.TaskId = tid
	err := ts.repo.CreateTask(task)
	if err != nil {
		return "", err
	}
	return tid, nil
}

func (ts *taskService) UpdateTask(task model.Task) error {
	return ts.repo.UpdateTask(task)
}

func (ts *taskService) DeleteTask(taskid string) error {
	return ts.repo.DeleteTask(taskid)
}
