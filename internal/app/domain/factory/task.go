package factory

import (
	"gotask/internal/app/adapter/postgresql/model"
	"gotask/internal/app/domain"
)

// Task factory for domain.Task
type Task struct{}

// ToEntity generates domain.Task from model.Task
func (tf Task) ToEntity(tm model.Task) domain.Task {
	return domain.Task{
		ID:        tm.ID,
		Title:     tm.Title,
		Message:   tm.Message,
		CreatedAt: tm.CreatedAt,
		UpdatedAt: tm.UpdatedAt,
	}
}

// ToModel generates a model.Task from a domain.Task
func (tf Task) ToModel(te domain.Task) model.Task {
	return model.Task{
		ID:        te.ID,
		Title:     te.Title,
		Message:   te.Message,
		CreatedAt: te.CreatedAt,
		UpdatedAt: te.UpdatedAt,
	}
}

// ToManyEntities generates many domain.Task from array of model.Task
func (tf Task) ToManyEntities(
	tm []model.Task,
) []domain.Task {
	tmLen := len(tm)
	tasks := make([]domain.Task, tmLen)

	for i := 0; i < tmLen; i++ {
		tasks[i] = tf.ToEntity(tm[i])
	}

	return tasks
}

// ToManyModels generates many model.Task from array of domain.Task
func (tf Task) ToManyModels(
	te []domain.Task,
) []model.Task {
	teLen := len(te)
	tasks := make([]model.Task, teLen)

	for i := 0; i < teLen; i++ {
		tasks[i] = tf.ToModel(te[i])
	}

	return tasks
}
