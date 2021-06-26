package main

import (
	"fmt"
	"html/template"
	"os"
)

var (
	CSV_FILE       = os.Getenv("CSV")
	SENDER         = os.Getenv("EMAIL_SENDER")
	PASSWORD       = os.Getenv("SENDER_PASSWORD")
	EMAIL_TEMPLATE = os.Getenv("MAIL_TEMPLATE")
	SMTP_HOST      = os.Getenv("SMTP_HOST")
	SMTP_PORT      = os.Getenv("SMTP_PORT")
)

func main() {
	var handler IHandler

	template, err := template.ParseFiles(EMAIL_TEMPLATE)
	if err != nil {
		fmt.Printf("Error parsing HTML template file. Error: %+v", err)
		return
	}
	fetcher := CsvEmailFetcher{CSV_FILE}
	sender := EmailSender{
		SENDER,
		PASSWORD,
		SMTP_HOST,
		SMTP_PORT,
		template,
	}
	validator := EmailValidator{}
	handler = &EmailHandler{&fetcher, &validator, &sender}
	handler.Handle()
	//fmt.Printf("Csv File: %s sender: %s, Password: %s, Email_Template: %s, smtp HOST: %s, SMTP_PORT: %s", CSV_FILE, SENDER, PASSWORD, EMAIL_TEMPLATE, SMTP_HOST, SMTP_PORT)
}
