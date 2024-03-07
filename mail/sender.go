package mail

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func SendMail(mailMessage *gomail.Message) error {
	mailDialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return mailDialer.DialAndSend(mailMessage)
}