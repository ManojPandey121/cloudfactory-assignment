package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvEmailFetcher struct {
	SourcePath string
}

type IEmailFetcher interface {
	Fetch() ([]Email, error)
}

func (csvf CsvEmailFetcher) Fetch() (emails []Email, err error) {
	file, err := os.Open(csvf.SourcePath)
	if err != nil {
		fmt.Printf("Error reading data from file %s. Error: %+v\n", csvf.SourcePath, err)
		return emails, err
	}

	rawEmails, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return emails, err
	}

	if len(rawEmails) == 0 {
		return emails, fmt.Errorf("No Emails provided in file %s", csvf.SourcePath)
	}

	for _, email := range rawEmails {
		emails = append(emails, Email{email[0]})
	}
	return
}
