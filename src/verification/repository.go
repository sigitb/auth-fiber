package verification

import (
	"base-fiber/app/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(verification models.Verification) (models.Verification, error)
	Find(code string, types string,id_user int) (models.Verification, error)
	UpdateStatus(id int) error
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(verification models.Verification) (models.Verification, error) {
	err := r.db.Create(&verification).Error

	if err != nil{
		return verification, err
	}

	return verification, nil
}

func (r *repository) Find(code string, types string, id_user int) (models.Verification, error) {
	var verification models.Verification
	err := r.db.Where("code = ?", code).Where("user_id = ?", id_user).Where("types = ?",types).Where("status = ?",0).Find(&verification).Error
	if err != nil{
		return verification, err
	}
	return verification, nil
}

func (r *repository) UpdateStatus(id int) error {
	err := r.db.Model(&models.Verification{}).Where("id = ?", id).Update("status", 1).Error

	if err != nil{
		return err
	}
	return nil
}