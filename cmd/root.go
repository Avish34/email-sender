package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/mail"
	"os"

	"github.com/Avish34/email-sender/types"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "email-sender",
	Short: `email-sender is a CLI tool for sending mails`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running validation..")
		senderEmail, _ := cmd.Flags().GetString("email-address")
		_, err := mail.ParseAddress(senderEmail)
		if err != nil {
			fmt.Println("Error %s", err)
			os.Exit(1)
		}
		filePath, _ := cmd.Flags().GetString("json-file")
		present, err := doesFileexists(filePath)
		if err != nil {
			fmt.Println("Error %s", err)
			os.Exit(1)
		}
		if !present {
			fmt.Println("file not present %s", filePath)
			os.Exit(1)
		}
	},
}

func Execute() (string, []types.EmailInfo)  {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	var emails []types.EmailInfo
	senderEmail, err := rootCmd.Flags().GetString("email-address")
	if(err != nil) {
		fmt.Println("error %v", err)
		os.Exit(1)
	}
	jsonFile , _ := rootCmd.Flags().GetString("json-file")
	file, err := os.Open(jsonFile)
	value, _ := ioutil.ReadAll(file)
	json.Unmarshal(value, &emails)
	return senderEmail, emails

}

func init() {
	fmt.Println("Init function")
	rootCmd.Flags().StringP("email-address", "", "", "specify the email address from which email need to be sent")
	rootCmd.Flags().StringP("json-file", "", "", "specify json file path used for sending email")
}