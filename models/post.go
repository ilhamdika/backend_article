package models

import (
	"time"
)

type Post struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"type:varchar(200);not null"`
	Content     string    `gorm:"type:text;not null"`
	Category    string    `gorm:"type:varchar(100);not null"`
	CreatedDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Status      string    `gorm:"type:enum('publish','draft','thrash');not null"`
}