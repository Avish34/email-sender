package types

import (
	"fmt"
	"log"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

type EmailService interface {
	SendEmails(string, []EmailInfo) error
}

// EmailJsonService is a service used for sending emails
type EmailJsonService struct {
	MailjetClient *mailjet.Client
}

func(e *EmailJsonService) SendEmails(senderEmail string, emails []EmailInfo) error {
	messages := e.convertIntoMessages(senderEmail, emails)
	res, err := e.MailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
	return nil
}

func(e *EmailJsonService) convertIntoMessages(senderEmail string, emails []EmailInfo) mailjet.MessagesV31 {
	var messagesInfo []mailjet.InfoMessagesV31
	for _, val := range emails {
		msg := mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
			  Email: senderEmail,
			},
			To: &mailjet.RecipientsV31{
			  mailjet.RecipientV31 {
				Email: val.Recipient,
			  },
			},
			Subject: val.Subject,
			TextPart: val.Body,
		}
		messagesInfo = append(messagesInfo, msg)
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	return messages
}

func NewEmailJsonService(PublicKey, PrivateKey string) EmailService {
	mailjetClient := mailjet.NewMailjetClient(PublicKey, PrivateKey)
	return &EmailJsonService{MailjetClient: mailjetClient}
}

// EmailInfo is object representing a single email
type EmailInfo struct {
	Recipient   string `json:"recipient,omitempty"`
	Subject     string `json:"subject,omitempty"`
	Body        string `json:"body,omitempty"`
}