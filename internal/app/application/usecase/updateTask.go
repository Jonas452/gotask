package usecase

import (
	"gotask/internal/app/domain"
	"gotask/internal/app/domain/repository"
)

// UpdateTask updates a task in the database
func UpdateTask(tr repository.ITask, id string, task domain.Task) (rTask domain.Task, err error) {
	return tr.Update(id, task)
}
