package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Task represents the model of Task
type Task struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	Title       string `gorm:"type:text;not null"`
	Description string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// BeforeCreate run before insert
func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()

	return
}
