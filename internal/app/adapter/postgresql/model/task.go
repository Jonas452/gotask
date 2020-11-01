package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Task represents the model of Task
type Task struct {
	ID      string `gorm:"type:uuid;primaryKey"`
	Title   string `gorm:"type:text;not null"`
	Message string `gorm:"type:text"`
}

// BeforeCreate run before insert
func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()

	return
}
