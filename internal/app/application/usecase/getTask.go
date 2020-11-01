package usecase

import (
	"gotask/internal/app/domain"
	"gotask/internal/app/domain/repository"
)

// GetTask is an usecase to get a task from an ID
func GetTask(tr repository.ITask, id string) domain.Task {
	return tr.Get(id)
}
