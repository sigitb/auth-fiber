package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null;type:varchar(100)"`
	Name string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	RoleId uint `json:"role_id" gorm:"not null"`
	Status uint `json:"status" gorm:"default:0"`
	Role Role `gorm:"foreignKey:RoleId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}