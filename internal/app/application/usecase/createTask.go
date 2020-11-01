package usecase

import (
	"gotask/internal/app/domain"
	"gotask/internal/app/domain/repository"
)

//CreateTask is a usecase to insert a task at the database
func CreateTask(tr repository.ITask, task domain.Task) (rTask domain.Task, err error) {
	return tr.Create(task)
}
