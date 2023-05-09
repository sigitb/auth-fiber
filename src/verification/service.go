package verification

import (
	"base-fiber/app/models"
	"errors"
	"time"
)

type service struct {
	repository Repository
}

type Service interface {
	Save(input InputVerification) (models.Verification, error)
	Find(code string, types string, id_user int) (models.Verification, error)
	UpdateStatus(id int) error
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Save(input InputVerification) (models.Verification, error) {
	var verification models.Verification

	verification.Code = input.Code
	verification.Types = input.Types
	verification.UserId = uint(input.UserId)
	verification.Expired = time.Now().Add(time.Hour * 1)

	newVerification, err  := s.repository.Save(verification)
	if err != nil {
		return verification, err
	}
	return newVerification, nil
}

func (s *service) Find(code string, types string,id_user int) (models.Verification, error) {
	verification := models.Verification{}
	findVerification, err := s.repository.Find(code, types, id_user)
	if err != nil {
		return verification, err
	}
	if findVerification.ID == 0 {
		return verification, errors.New("Verification not found")
	}
	return findVerification, nil
}

func (s *service) UpdateStatus(id int) error {
	err := s.repository.UpdateStatus(id)
	if err != nil {
		return err
	}

	return nil
}
