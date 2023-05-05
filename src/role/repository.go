package role

import (
	"base-fiber/app/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindById(id int) (models.Role, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindById(id int) (models.Role, error) {
	var role models.Role
	err := r.db.Where("id = ?", id).Find(&role).Error
	if err != nil{
		return role, err
	}
	return role, nil
}