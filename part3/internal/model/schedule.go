package model

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	TaskID    uint           `gorm:"not null;index" json:"task_id"`
	StartAt   time.Time      `gorm:"not null" json:"start_at"`
	EndAt     time.Time      `gorm:"not null" json:"end_at"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Task      Task           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
