package domain

// Task represents the tasks in the system
type Task struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}
