package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"squad10x.com.br/boilerplate/internal/pkg/db/repositories"
	"squad10x.com.br/boilerplate/internal/pkg/dtos"
	"squad10x.com.br/boilerplate/internal/pkg/enums"
	m "squad10x.com.br/boilerplate/internal/pkg/models"
	"squad10x.com.br/boilerplate/internal/pkg/notifications"
	"squad10x.com.br/boilerplate/internal/pkg/services"
	"squad10x.com.br/boilerplate/internal/pkg/utils"
	"squad10x.com.br/boilerplate/pkg/logger"
)

type userHandler struct {
	logger         logger.Logger
	userRepository repositories.UserRepository
	userService    services.UserService
	notification   notifications.NotificationService
}

func UserHandler() *userHandler {
	return &userHandler{
		logger:         logger.GetLogger(),
		userRepository: repositories.NewUserRepository(),
		userService:    services.NewUserService(),
		notification:   notifications.NewNotificationService(),
	}
}

func (h *userHandler) Create(u m.User) (dtos.UserDto, error) {
	u.RoleId = int(enums.User)
	u = h.userRepository.Create(u)
	return u.ParseDto(), nil
}

func (h *userHandler) ForgotPassword(email string) (dtos.UserDto, error) {
	user, _ := h.userRepository.GetByEmail(email, utils.Join{
		Rels: []string{"Confirmation"},
	})

	user.Confirmation.Confirmed = false
	user.Confirmation.Code = h.userService.GenerateCode()
	user.Confirmation.UserId = user.ID

	h.userRepository.SaveConfirmation(&user.Confirmation)
	err := h.notification.SendForgotPassword(user)
	if err != nil {
		return dtos.UserDto{}, err
	}

	return user.ParseDto(), nil
}

func (h *userHandler) ValidateCode(email, code string) (fiber.Map, error) {
	user, _ := h.userRepository.GetByEmail(email, utils.Join{Rels: []string{"Confirmation"}})

	res, err := h.userService.ValidateCode(code, user)
	if err != nil {
		return fiber.Map{}, err
	}

	return fiber.Map{"status": res}, nil
}

func (h *userHandler) ChangePassword(code, password, email string) (dtos.UserDto, error) {
	user, _ := h.userRepository.GetByEmail(email, utils.Join{Rels: []string{"Confirmation"}})

	codeValid, err := h.userService.ValidateCode(code, user)
	if err != nil || !codeValid {
		return dtos.UserDto{}, err
	}

	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return dtos.UserDto{}, err
	}

	user.Confirmation.Confirmed = true
	user.Password = string(encrypted)

	h.userRepository.Save(&user)
	h.userRepository.SaveConfirmation(&user.Confirmation)

	return user.ParseDto(), nil
}
