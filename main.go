package main

import (
	"api/internal/config"
	"api/internal/controller"
	"api/internal/database"
	"api/internal/middleware"
	"api/internal/repository"
	"api/internal/service"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Start() error {

	dbConfig := config.NewConfig()
	db, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	taskRepo := repository.NewTaskRepo(db)
	taskService := service.NewService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	router := http.NewServeMux()

	router.HandleFunc("GET /tasks", taskHandler.GetAllTasks)
	router.HandleFunc("POST /tasks", taskHandler.CreateTask)
	router.HandleFunc("GET /task/{id}", taskHandler.GetTask)
	router.HandleFunc("PUT /task/{id}", taskHandler.UpdateTask)
	router.HandleFunc("DELETE /task/{id}", taskHandler.DeleteTask)

	middlewareChain := middleware.MiddlewareChain(middleware.Logger, middleware.AuthMiddleware)

	server := &http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}

	log.Printf("Starting API server on %s", s.addr)

	return server.ListenAndServe()
}

func main() {
	server := NewAPIServer(config.NewConfig().Port)
	server.Start()
}
