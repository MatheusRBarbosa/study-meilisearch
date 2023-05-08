package notifications

import (
	i "squad10x.com.br/boilerplate/domain/interfaces"
	m "squad10x.com.br/boilerplate/domain/models"
	"squad10x.com.br/boilerplate/infra/mailer"
	"squad10x.com.br/boilerplate/infra/mailer/mailables"
)

type notificationService struct {
	mailService i.MailService
}

func NotificationService() i.NotificationService {
	return &notificationService{
		mailService: mailer.MailService(),
	}
}

func (n *notificationService) SendForgotPassword(u m.User) error {
	var forgot mailables.ForgotPasswordMail
	mail := forgot.Prepare(u.Email, u)

	go n.mailService.Send(mail)

	return nil
}
