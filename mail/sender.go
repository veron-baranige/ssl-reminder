package mail

import (
	"crypto/tls"

	"github.com/veron-baranige/ssl-reminder/config"
	"gopkg.in/gomail.v2"
)

func SendMail(subject string, content string) error {
	mailDialer := gomail.NewDialer(
		config.SmtpHost,
		config.SmtpPort,
		config.SmtpUser, 
		config.SmtpPassword,
	)

	message := gomail.NewMessage()
	message.SetHeader("From", config.SmtpUser)
	message.SetHeader("To", config.ReceiverMailAddresses...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", content)
	
	mailDialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return mailDialer.DialAndSend(message)
}