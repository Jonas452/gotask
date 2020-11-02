package domain

import "time"

// Task represents the tasks in the system
type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title" binding:"required,min=2,max=150"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
