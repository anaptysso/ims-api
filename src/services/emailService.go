package services

import (
	config "imsapi/config"
	enum "imsapi/src/enum"
	"net/smtp"
	"strings"
)

// EmailService provides the way to interact with 3rd party through email
type EmailService struct {
	Config *config.Configuration
}

// SendMail function sends email to other email address
func (es *EmailService) SendMail(to []string, subject string, body string) {
	if es.Config.SMTP.EnableService == enum.ServiceEnabled {
		message := []byte("To: " + strings.Join(to, ",") + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" + body)

		auth := smtp.PlainAuth("", es.Config.SMTP.UserName, es.Config.SMTP.Password, es.Config.SMTP.Host)
		err := smtp.SendMail(es.Config.SMTP.Host+":"+es.Config.SMTP.Port, auth, es.Config.SMTP.UserName, to, message)

		if err != nil {
			panic(err)
		}
	} else {
		panic("Email service isn't enabled")
	}
}
