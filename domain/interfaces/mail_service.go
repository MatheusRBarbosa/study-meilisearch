package interfaces

import m "squad10x.com.br/boilerplate/domain/models"

type MailService interface {
	Send(m.Mailable) error
}
