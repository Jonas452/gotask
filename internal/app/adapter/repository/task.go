package repository

import (
	"gotask/internal/app/adapter/postgresql"
	"gotask/internal/app/adapter/postgresql/model"
	"gotask/internal/app/domain"
	"gotask/internal/app/domain/factory"
)

// Task represents the repository for domain.Task
type Task struct{}

// Get gets a task based on an id
func (t Task) Get(id string) domain.Task {
	db := postgresql.Connection()

	var mTask model.Task
	result := db.Where("id = ?", id).Find(&mTask)

	if result.Error != nil {
		panic(result.Error)
	}

	taskFactory := factory.Task{}
	return taskFactory.ToEntity(mTask)
}

// GetAll gets all tasks from database
func (t Task) GetAll() []domain.Task {
	db := postgresql.Connection()

	var mTasks []model.Task
	result := db.Find(&mTasks)
	if result.Error != nil {
		panic(result.Error)
	}

	taskFactory := factory.Task{}
	return taskFactory.ToManyEntities(mTasks)
}

// Create inserts a task in the database
func (t Task) Create(eTask domain.Task) (rTask domain.Task, err error) {
	db := postgresql.Connection()

	taskFactory := factory.Task{}
	mTask := taskFactory.ToModel(eTask)

	result := db.Create(&mTask)
	if result.Error != nil {
		return domain.Task{}, result.Error
	}

	return taskFactory.ToEntity(mTask), nil
}

// Update updates a task in the database
func (t Task) Update(id string, eTask domain.Task) (rTasl domain.Task, err error) {
	db := postgresql.Connection()

	taskFactory := factory.Task{}
	mTask := taskFactory.ToModel(eTask)
	mTask.ID = id

	result := db.Save(&mTask)
	if result.Error != nil {
		return domain.Task{}, result.Error
	}

	return taskFactory.ToEntity(mTask), nil
}

// Delete deletes a task from the database
func (t Task) Delete(id string) (err error) {
	db := postgresql.Connection()

	var mTask model.Task
	result := db.Where("id = ?", id).Delete(&mTask)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
