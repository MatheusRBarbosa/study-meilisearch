package mailables

import m "squad10x.com.br/boilerplate/internal/pkg/models"

type ForgotPasswordMail m.Mailable

func (f *ForgotPasswordMail) Prepare(to string, model interface{}) m.Mailable {
	return m.Mailable{
		To:           []string{to},
		Model:        model,
		Subject:      "Recuperação de senha",
		TemplateName: "forgot_password.html",
		ResorcesName: []string{"logo.png"},
	}
}
