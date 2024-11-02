package handler

import (
	"api/internal/model"
	"api/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type taskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *taskHandler {
	return &taskHandler{service: service}
}

func (th *taskHandler) GetAllTasks(w http.ResponseWriter, _ *http.Request) {
	tasks, err := th.service.GetAllTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"success": true, "tasks": tasks})
}

func (th *taskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := th.service.CreateTask(task); err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"success": true})
}

func (th *taskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := th.service.UpdateTask(task); err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"success": true})
}

func (th *taskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	fmt.Println(id)

	// if err := th.service.DeleteTask(id); err != nil {
	// 	http.Error(w, fmt.Sprintf("Failed to delete task with ID %d", id), http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"success": true})
}