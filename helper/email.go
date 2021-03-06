package helper

import (
	"fmt"
	"net/smtp"
)

const (
	fromAddress       = "pranotobudi.app@gmail.com"
	fromEmailPassword = "CamelCasePasswordBismillah"
	smtpServer        = "smtp.gmail.com"
	smtpPort          = "587"
)

func SendEmail(toAddress []string, message []byte) {
	// Message.
	//   message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", fromAddress, fromEmailPassword, smtpServer)

	// Sending email.
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, fromAddress, toAddress, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
