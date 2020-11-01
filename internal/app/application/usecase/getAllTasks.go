package usecase

import (
	"gotask/internal/app/domain"
	"gotask/internal/app/domain/repository"
)

//GetAllTasks is a use case to get all tasks
func GetAllTasks(tr repository.ITask) []domain.Task {
	return tr.GetAll()
}
