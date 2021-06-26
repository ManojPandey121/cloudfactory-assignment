package main

import (
	"bytes"
	"html/template"
	"net/smtp"
)

type EmailSender struct {
	Sender         string
	SenderPassword string
	SmtpHost       string
	SmtpPort       string
	EmailTemplate  *template.Template
}

type IEmailSender interface {
	Send(Email) error
}

func (es EmailSender) Send(email Email) error {
	var body bytes.Buffer
	// Authentication.
	auth := smtp.PlainAuth("", es.Sender, es.SenderPassword, es.SmtpHost)
	es.EmailTemplate.Execute(&body, email)
	return smtp.SendMail(es.SmtpHost+":"+es.SmtpPort, auth, es.Sender, []string{email.Value}, body.Bytes())
}
