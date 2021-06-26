package main

import (
	"regexp"
)

type EmailValidator struct{}

type IEmailValidator interface {
	Validate(Email) bool
}

const EMAIL_REGEX = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

func (ev EmailValidator) Validate(email Email) bool {
	re := regexp.MustCompile(EMAIL_REGEX)
	return re.MatchString(email.Value)
}
