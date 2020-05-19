package services

import (
	"fmt"
	"net/smtp"
	"strings"

	enableServiceEnum "imsapi/enums/enableService"
	environmentEnum "imsapi/enums/environment"
)

type EmailService struct {
	Service
}

func (this EmailService) SendMailSystem(to []string, subject string, body string) {
	if this.Configuration.Smtp.EnableService == enableServiceEnum.Yes {
		message := []byte("To: " + strings.Join(to, ",") + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" + body)

		auth := smtp.PlainAuth("", this.Configuration.Smtp.UserName, this.Configuration.Smtp.Password, this.Configuration.Smtp.Host)
		err := smtp.SendMail(this.Configuration.Smtp.Host+":"+this.Configuration.Smtp.Port, auth, this.Configuration.Smtp.UserName, to, message)

		if err != nil {
			panic(err)
		}
	} else {
		if this.Configuration.Environment == environmentEnum.Live {
			panic("Email service isn't enabled")
		} else if this.Configuration.Environment == environmentEnum.Development {
			fmt.Println("Email service isn't enabled")
		}
	}
}
