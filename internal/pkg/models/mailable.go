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
	ResorcesName []string
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

func (m *Mailable) GetResourcesPath() []string {
	if len(m.ResorcesName) == 0 {
		return []string{}
	}

	pwd, _ := os.Getwd()
	base := pwd + "/infra/mailer/resources/"

	var resources []string

	for i := 0; i < len(m.ResorcesName); i++ {
		fileName := m.ResorcesName[i]
		resources = append(resources, base+fileName)
	}

	return resources
}

func getTemplatePath(templateName string) string {
	pwd, _ := os.Getwd()
	path := pwd + "/infra/mailer/templates/" + templateName

	fmt.Printf(path)
	return path
}
