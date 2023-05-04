package mailer

import (
	"fmt"

	i "squad10x.com.br/boilerplate/domain/interfaces"
	m "squad10x.com.br/boilerplate/domain/models"
)

type mailService struct{}

func MailService() i.MailService {
	return &mailService{}
}

func (s *mailService) Send(mail m.Mailable) {
	fmt.Println(mail)
}
