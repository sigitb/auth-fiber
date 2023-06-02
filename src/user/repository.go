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
	UpdateStatus(email string) error
	UpdatePassword(password string) error
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
func (r *repository) UpdateStatus(email string) error {
	err := r.db.Model(&models.User{}).Where("email = ?", email).Where("deleted_at IS NULL").Update("status", 1).Error
	if err != nil{
		return err
	}
	return nil
}

func (r *repository) UpdatePassword(password string) error {
	err := r.db.Model(&models.User{}).Where("password = ?", password).Where("deleted_at IS NULL").Update("status", 1).Error
	if err != nil{
		return err
	}
	return nil
}