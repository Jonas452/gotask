package usecase

import (
	"gotask/internal/app/domain"
	"gotask/internal/app/domain/repository"
)

// UpdateTask updates a task in the database
func UpdateTask(tr repository.ITask, task domain.Task) (domain.Task, error) {
	return tr.Update(task)
}
