package models

import (
	"time"

	"gorm.io/gorm"
)

type Verification struct {
	gorm.Model
	Code string `json:"code" gorm:"type:varchar(20)"`
	UserId uint `json:"user_id" gorm:"not null"`
	Types string `json:"types" gorm:"type:varchar(50)"`
	Status uint `json:"status" gorm:"default:0"`
	User User `gorm:"foreignKey:UserId"`
	Expired time.Time `json:"expired" gorm:"not null"`
}