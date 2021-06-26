package main

import (
	"errors"
	"fmt"
	"testing"
)

type mockEmailFetcher struct {
	Emails []Email
	err    error
}
type mockValidator struct {
	validEmail bool
}
type mockSender struct {
	err error
}

func (m mockEmailFetcher) Fetch() ([]Email, error) {
	return m.Emails, m.err
}

func (m mockValidator) Validate(email Email) bool {
	return m.validEmail
}

func (m mockSender) Send(email Email) error {
	return m.err
}

func TestEmailHandler_Handle(t *testing.T) {
	t.Log("When there is error fetching emails")
	{
		t.Run("logs error without crashing and doesnot return anything", func(t *testing.T) {
			subject := EmailHandler{
				EmailFetcher: &mockEmailFetcher{[]Email{}, errors.New("Error fetching email")},
				Validator:    &mockValidator{},
				Sender:       &mockSender{},
			}
			subject.Handle()
		})
	}
	t.Log("When there is error sending emails")
	{
		t.Run("logs error without crashing and doesnot return anything", func(t *testing.T) {
			subject := EmailHandler{
				EmailFetcher: &mockEmailFetcher{[]Email{Email{"manoj.pandey@gmail.com"}}, nil},
				Validator:    &mockValidator{},
				Sender:       &mockSender{errors.New("Error sending email")},
			}
			subject.Handle()
		})
	}
	t.Log("When everything is fine")
	{
		t.Run("sends email and doesnot return anything", func(t *testing.T) {
			subject := EmailHandler{
				EmailFetcher: &mockEmailFetcher{[]Email{Email{"manoj.pandey@gmail.com"}}, nil},
				Validator:    &mockValidator{},
				Sender:       &mockSender{},
			}
			subject.Handle()
		})
	}
}

func TestEmailHandler_ValidateAndSendEmail(t *testing.T) {
	t.Log("When email is not valid")
	{
		t.Run("returns error", func(t *testing.T) {
			email := "invalid_email"
			subject := EmailHandler{
				Validator: &mockValidator{false},
				Sender:    &mockSender{},
			}
			err := subject.validateAndSendEmail(Email{email})
			if err == nil {
				t.Fatalf("Expected error got nil instead. Error: %+v", err)
			}
			expErrorMsg := fmt.Sprintf("Email: %s is not a valid email", email)
			if err.Error() != expErrorMsg {
				t.Fatalf("Expected error message %s, got %s", expErrorMsg, err.Error())
			}
		})
	}
	t.Log("When there was error sending mail")
	{
		t.Run("returns error", func(t *testing.T) {
			email := "email"
			subject := EmailHandler{
				Validator: &mockValidator{true},
				Sender:    &mockSender{errors.New("Error sending email")},
			}
			err := subject.validateAndSendEmail(Email{email})
			if err == nil {
				t.Fatalf("Expected error got nil instead. Error: %+v", err)
			}
			expErrorMsg := "Error sending email"
			if err.Error() != expErrorMsg {
				t.Fatalf("Expected error message %s, got %s", expErrorMsg, err.Error())
			}
		})
	}
}
