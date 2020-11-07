package domain

import "time"

// Task represents the tasks in the system
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" binding:"required,min=2,max=150"`
	Description string    `json:"description" bindind:"omitempty,min=2,max=300"`
	DueDate     time.Time `json:"due_date" binding:"omitempty,gt"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
