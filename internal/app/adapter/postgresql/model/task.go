package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Task represents the model of Task
type Task struct {
	ID          string         `gorm:"type:uuid;primaryKey"`
	Title       string         `gorm:"type:text;not null"`
	Description sql.NullString `gorm:"type:text;default:null"`
	DueDate     sql.NullTime   `gorm:"default:null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// BeforeCreate run before insert
func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()

	return
}
