package usecase

import "gotask/internal/app/domain/repository"

// DeleteTask deletes a task from the database
func DeleteTask(tr repository.ITask, id string) (err error) {
	return tr.Delete(id)
}
