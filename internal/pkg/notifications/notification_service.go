package notifications

import (
	"squad10x.com.br/boilerplate/internal/pkg/mailer"
	"squad10x.com.br/boilerplate/internal/pkg/mailer/mailables"
	m "squad10x.com.br/boilerplate/internal/pkg/models"
)

type NotificationService interface {
	SendForgotPassword(m.User) error
}

type notificationService struct {
	mailService mailer.MailService
}

func NewNotificationService() NotificationService {
	return &notificationService{
		mailService: mailer.NewMailService(),
	}
}

func (n *notificationService) SendForgotPassword(u m.User) error {
	var forgot mailables.ForgotPasswordMail
	mail := forgot.Prepare(u.Email, u)

	go n.mailService.Send(mail)

	return nil
}
