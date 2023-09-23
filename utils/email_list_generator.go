package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Avish34/email-sender/types"
)

func writeIntoJsonFile(fileName string, emails []types.EmailInfo) error {
	emailsJson, _ := json.Marshal(emails)
    err := ioutil.WriteFile(fileName, emailsJson, 0644)
    fmt.Println("%+v", emailsJson)
	return err
}

func Generate_emails(totalEmails int, emails types.EmailInfo, fileName string) error {
	if totalEmails == 0 {
		fmt.Printf("Email genreation can't be zero")
		return nil
	}
	var listOfEmails []types.EmailInfo
	for totalEmails > 0 {
		listOfEmails = append(listOfEmails, emails)
		totalEmails--
	}
	return writeIntoJsonFile(fileName, listOfEmails)
}