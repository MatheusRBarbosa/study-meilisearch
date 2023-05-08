package interfaces

import m "squad10x.com.br/boilerplate/domain/models"

type NotificationService interface {
	SendForgotPassword(m.User) error
}
