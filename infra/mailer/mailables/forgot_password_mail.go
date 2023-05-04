package mailables

import m "squad10x.com.br/boilerplate/domain/models"

type ForgotPasswordMail m.Mailable

func (f *ForgotPasswordMail) Prepare(to string, model interface{}) m.Mailable {
	return m.Mailable{
		To:           to,
		Model:        model,
		Subject:      "Recuperação de senha",
		TemplatePath: "",
	}
}
