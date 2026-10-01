package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/artemSoko1ov/LifeOS-backend/internal/database"
	"github.com/artemSoko1ov/LifeOS-backend/internal/task"
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
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env")
	}

	db, err := database.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("database connected")

	mux := http.NewServeMux()

	taskStorage := &task.Storage{
		Tasks: []task.Task{},
	}

	taskRepository := task.NewRepository(taskStorage)
	taskService := task.NewService(taskRepository)
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
