package user

import (
	"base-fiber/app/models"
	"base-fiber/app/utils"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type service struct {
	repository Repository
}

type Service interface {
	RegisterUser(input InputRegister) (models.User, error)
	Login(input InputLogin) (string, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}


func (s *service) RegisterUser(input InputRegister) (models.User, error) {
	user := models.User{}
	
	checkEmail, err := s.repository.FindByEmail(input.Email)
	if err != nil{
		return user, err
	}
	
	if checkEmail.ID != 0{
		return user, errors.New("Email already used")
	}
	
	user.Username = input.Username
	user.Email = input.Email
	user.Name = input.Name
	user.RoleId = input.RoleId
	passwordHash, err := utils.HashingPassword(input.Password)
	if err != nil{
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil{
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input InputLogin) (string, error) {
	checkEmail, err := s.repository.FindByEmail(input.Email)
	if err != nil{
		return "", err
	}
	
	if checkEmail.ID == 0{
		return "", errors.New("User not found")
	}

	if !utils.CheckPasswordHash(input.Password, checkEmail.Password) {
		return "" , errors.New("Password is wrong")
	}

	claims := jwt.MapClaims{}
	claims["name"] = checkEmail.Name
	claims["email"] = checkEmail.Email
	claims["username"] = checkEmail.Username
	claims["role_id"] = checkEmail.RoleId
	claims["exp"] = time.Now().Add(time.Minute * 24).Unix()

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		return "" , errGenerateToken
	}

	return token, nil 
}