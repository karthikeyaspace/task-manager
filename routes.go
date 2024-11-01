package main

import (
	"api/internal/middleware"
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
	router := http.NewServeMux()

	// routes

	router.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("get tasks"))
	})

	router.HandleFunc("POST /tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("post tasks"))
	})

	router.HandleFunc("GET /task/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("get task by id" + r.URL.Query().Get("id")))
	})

	router.HandleFunc("PUT /task/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("put task by id" + r.URL.Query().Get("id")))
	})

	router.HandleFunc("DELETE /task/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("delete task by id" + r.URL.Query().Get("id")))
	})

	// middleware is on every route
	middlewareChain := middleware.MiddlewareChain(middleware.Logger, middleware.AuthMiddleware)

	server := &http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}

	log.Printf("Starting API server on %s", s.addr)

	return server.ListenAndServe()
}
