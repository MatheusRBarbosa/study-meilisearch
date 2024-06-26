package mailer

import (
	"crypto/tls"
	"strconv"
	"strings"

	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	m "squad10x.com.br/boilerplate/internal/pkg/models"
	"squad10x.com.br/boilerplate/pkg/configs"
)

type MailService interface {
	Send(m.Mailable) error
}

type smtpConfig struct {
	Host     string
	Port     int
	From     string
	Username string
	Password string
}

type mailService struct {
	Config *smtpConfig
}

var config *smtpConfig

func init() {
	port, _ := strconv.Atoi(configs.GetEnv("SMTP_PORT"))
	config = &smtpConfig{
		Host:     configs.GetEnv("SMTP_HOST"),
		Port:     port,
		From:     configs.GetEnv("SMTP_FROM"),
		Username: configs.GetEnv("SMTP_USERNAME"),
		Password: configs.GetEnv("SMTP_PASSWORD"),
	}
}

func NewMailService() MailService {
	return &mailService{
		Config: config,
	}
}

func (s *mailService) Send(mail m.Mailable) error {
	msg, err := s.createMessage(mail)
	if err != nil {
		return err
	}

	dialer := gomail.NewDialer(s.Config.Host, s.Config.Port, s.Config.Username, s.Config.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}

func (s *mailService) createMessage(mail m.Mailable) (*gomail.Message, error) {
	data, err := mail.ParseTemplate()
	if err != nil {
		return nil, err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", s.Config.From)
	msg.SetHeader("To", strings.Join(mail.To, ", "))
	msg.SetHeader("Subject", mail.Subject)
	msg.SetBody("text/html", data)
	msg.AddAlternative("text/plain", html2text.HTML2Text(data))

	resources := mail.GetResourcesPath()
	for i := 0; i < len(resources); i++ {
		r := resources[i]
		msg.Embed(r)
	}

	return msg, nil
}
