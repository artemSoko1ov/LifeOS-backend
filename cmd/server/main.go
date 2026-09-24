package main

import (
	"github.com/artemSoko1ov/LifeOS-backend/internal/task"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	taskStorage := &task.Storage{
		Tasks: []task.Task{
			{
				ID:        "1",
				Title:     "Изучить Go",
				Completed: false,
			},
		},
	}
	taskHandler := task.NewHandler(taskStorage)

	mux.HandleFunc("GET /api/tasks", taskHandler.GetTasks)
	mux.HandleFunc("POST /api/tasks", taskHandler.CreateTask)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("server started on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
