package services

type EmailService interface {
	SendMailSystem(to []string, subject string, body string)
}
