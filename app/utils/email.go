package utils

import (
	"base-fiber/app/utils/templates"
	"net/url"
	"os"

	"gopkg.in/gomail.v2"
)

func ForgotPassword(token string, email string, name string) {
	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("EMAIL_FROM_ADDRESS"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Forgot password!")
	m.SetBody("text/html", templates.TemplateResetPassword(os.Getenv("BASE_URL")+"?token="+url.QueryEscape(token)+"&email="+url.QueryEscape(email), name))

	d := gomail.NewDialer(
		os.Getenv("EMAIL_HOST"),
		StringToInt(os.Getenv("EMAIL_PORT")),
		os.Getenv("EMAIL_USERNAME"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func Verification(name string, email string, token string, url string) {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_FROM_ADDRESS"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email Verifikasi!")
	m.SetBody("text/html", templates.TempalteVerification(name, email, token, url))

	d := gomail.NewDialer(
		os.Getenv("EMAIL_HOST"),
		StringToInt(os.Getenv("EMAIL_PORT")),
		os.Getenv("EMAIL_USERNAME"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}