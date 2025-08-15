package entities

import (
	"time"
)

type Article struct {
	Id        int    `gorm:"primaryKey"`
	Title     string `gorm:"size:120;not null"`
	Type      string `gorm:"size:32;not null;default:html"`
	Author    string `gorm:"size:32;default:'Mr.Blue'"`
	Abstract  string `gorm:"size:255;"`
	Content   string `gorm:"type:text;not null"` // html内容
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string `gorm:"size:20;default:'draft'"` // draft/published/deleted
	Category  string `gorm:"size:20;default:'default'"`
	Tags      string `gorm:"size:40;"`
	//CoverImage string `gorm:"size:255"`                // 封面图URL
}
