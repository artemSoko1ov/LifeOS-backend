package main

import (
	"github.com/artemSoko1ov/LifeOS-backend/internal/task"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	taskStorage := &task.Storage{}
	taskHandler := task.NewHandler(taskStorage)

	mux.HandleFunc("GET /api/tasks", taskHandler.GetTasks)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("server started on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
