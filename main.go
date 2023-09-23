package main

import (
	"fmt"
	"os"

	"github.com/Avish34/email-sender/cmd"
	"github.com/Avish34/email-sender/types"
	"github.com/Avish34/email-sender/utils"
)

func main() {

	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	privateKey := os.Getenv("MJ_APIKEY_PRIVATE")
	dummyUser := types.EmailInfo{
		Recipient: "bapnadivisha@gmail.com",
		Body: "Hi, good night",
		Subject: "Congrats! Open the mail for more info",
	}
	fmt.Print("Generating dummy emails")
	utils.Generate_emails(5, dummyUser, "output1.json")
	senderService := types.NewEmailJsonService(publicKey, privateKey)
	senderEmail, emails := cmd.Execute()
	senderService.SendEmails(senderEmail, emails)
	fmt.Printf("senderEmail %s, totalEmails %d", senderEmail, len(emails))
}