package main

import "api/internal/config"

func main() {
	server := NewAPIServer(config.NewConfig().Port)
	server.Start()
}
