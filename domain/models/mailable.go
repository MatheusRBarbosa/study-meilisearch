package models

type Mailable struct {
	Subject      string
	To           string
	TemplatePath string
	Model        interface{}
}
