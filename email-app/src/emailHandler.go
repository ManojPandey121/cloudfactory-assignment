package main

import (
	"fmt"
	"sync"
)

type EmailHandler struct {
	EmailFetcher IEmailFetcher
	Validator    IEmailValidator
	Sender       IEmailSender
}

type IHandler interface {
	Handle()
}

func (eh *EmailHandler) Handle() {

	emails, err := eh.EmailFetcher.Fetch()
	if err != nil {
		fmt.Printf("Error fetching Users emails. Error: %+v", err)
	}

	var wg sync.WaitGroup
	for _, email := range emails {
		wg.Add(1)
		go func(email Email) {
			defer wg.Done()

			err := eh.validateAndSendEmail(email)
			if err != nil {
				fmt.Printf("Error sending email to %s. Error: %v.\n", email.Value, err)
			}
			return
		}(email)
	}
	// wait until all goroutines complete
	wg.Wait()
}

func (eh *EmailHandler) validateAndSendEmail(email Email) error {
	if !eh.Validator.Validate(email) {
		errorMsg := fmt.Sprintf("Email: %s is not a valid email", email.Value)
		return fmt.Errorf("%s", errorMsg)
	}

	err := eh.Sender.Send(email)
	if err == nil {
		fmt.Printf("Mail successfully send to %s.\n", email.Value)
	}
	return err
}
