package user

import (
	"base-fiber/app/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	if err != nil{
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil{
		return user, err
	}
	return user, nil
}