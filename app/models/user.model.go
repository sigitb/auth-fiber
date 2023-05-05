package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Name string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	RoleId uint `json:"role_id"`
	Role Role `gorm:"foreignKey:RoleId"`
}