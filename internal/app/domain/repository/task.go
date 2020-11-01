package repository

import "gotask/internal/app/domain"

// ITask is an interface of Task repository
type ITask interface {
	Get(id string) domain.Task
	GetAll() []domain.Task
	Create(eTask domain.Task) (task domain.Task, err error)
	Update(id string, eTask domain.Task) (task domain.Task, err error)
	Delete(id string) (err error)
}
