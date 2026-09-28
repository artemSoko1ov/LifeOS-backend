package main

import (
	"github.com/artemSoko1ov/LifeOS-backend/internal/task"
	"log"
	"net/http"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	taskStorage := &task.Storage{}
	taskService := task.NewService(taskStorage)
	taskHandler := task.NewHandler(taskService)

	mux.HandleFunc("GET /api/tasks", taskHandler.GetTasks)
	mux.HandleFunc("POST /api/tasks", taskHandler.CreateTask)

	server := http.Server{
		Addr:    ":8080",
		Handler: enableCORS(mux),
	}

	log.Println("server started on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
