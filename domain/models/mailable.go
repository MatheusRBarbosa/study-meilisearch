package models

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

type Mailable struct {
	Subject      string
	To           []string
	TemplateName string
	Model        interface{}
}

func (m *Mailable) ParseTemplate() (string, error) {
	path := getTemplatePath(m.TemplateName)
	t, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, m.Model); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func getTemplatePath(templateName string) string {
	pwd, _ := os.Getwd()
	path := pwd + "/infra/mailer/templates/" + templateName

	fmt.Printf(path)
	return path
}
