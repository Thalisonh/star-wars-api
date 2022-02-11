package models

import (
	"time"

	"gorm.io/gorm"
)

type Planets struct {
	gorm.Model
	CreatedAt  *time.Time `gorm:"<-:create"`
	Name       string     `gorm:"name" json:"name"`
	Clime      string     `gorm:"clime" json:"clime"`
	Ground     string     `gorm:"ground" json:"ground"`
	FilmsCount int        `gorm:"filmsCount" json:"filmsCount"`
}
