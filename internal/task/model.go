package task

import "time"

type Task struct {
	ID        string    `json:"id"`
	UserID    string	`json:"-"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
