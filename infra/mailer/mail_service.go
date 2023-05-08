package mailer

import (
	"crypto/tls"
	"strconv"
	"strings"

	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"squad10x.com.br/boilerplate/crosscutting"
	i "squad10x.com.br/boilerplate/domain/interfaces"
	m "squad10x.com.br/boilerplate/domain/models"
)

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
	port, _ := strconv.Atoi(crosscutting.GetEnv("SMTP_PORT"))
	config = &smtpConfig{
		Host:     crosscutting.GetEnv("SMTP_HOST"),
		Port:     port,
		From:     crosscutting.GetEnv("SMTP_FROM"),
		Username: crosscutting.GetEnv("SMTP_USERNAME"),
		Password: crosscutting.GetEnv("SMTP_PASSWORD"),
	}
}

func MailService() i.MailService {
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

	return msg, nil
}
