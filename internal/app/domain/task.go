package domain

import (
	"encoding/json"
	"time"
)

// Task represents the tasks in the system
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" binding:"required,min=2,max=150"`
	Description string    `json:"description" bindind:"omitempty,min=2,max=300"`
	DueDate     time.Time `json:"due_date" binding:"omitempty,gt"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MarshalJSON custom marshal for Task
func (t *Task) MarshalJSON() ([]byte, error) {
	type Alias Task
	var dueDate string
	if !t.DueDate.IsZero() {
		dueDate = t.DueDate.String()
	}

	return json.Marshal(&struct {
		DueDate string `json:"due_date"`
		*Alias
	}{
		DueDate: dueDate,
		Alias:   (*Alias)(t),
	})

}
