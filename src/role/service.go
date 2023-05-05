package role

import (
	"base-fiber/app/models"
	"errors"
)

type service struct {
	repository Repository
}

type Service interface {
	FindById(id int) (models.Role, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindById(id int) (models.Role, error) {
	role := models.Role{}
	dateRole, err := s.repository.FindById(id)
	if err != nil{
		return role, err 
	}
	if dateRole.ID == 0 {
		return role, errors.New("Role not found")
	}
	return dateRole, nil
}