package user

import "base-fiber/app/models"

type UserFormatter struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type LoginFormatter struct {
  Token string `json:"token"`
}

func FormatUser(user models.User) UserFormatter {
	formatter := UserFormatter{
		Id:       int(user.ID),
		Email:    user.Email,
		Username: user.Username,
		Name:     user.Name,
	}
	return formatter
}

func FormaterLogin(login LoginFormatter) LoginFormatter {
	formmater := LoginFormatter{
		Token: login.Token,
	}
	return formmater
}

