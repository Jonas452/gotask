package repository

import "gotask/internal/app/domain"

// ITask is an interface of Task repository
type ITask interface {
	Get(id string) domain.Task
	GetAll() []domain.Task
	Create(eTask domain.Task) (domain.Task, error)
	Update(eTask domain.Task) (domain.Task, error)
	Delete(id string) error
}
