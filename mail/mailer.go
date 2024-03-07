package mail

import (
	"github.com/veron-baranige/ssl-reminder/config"
	"gopkg.in/gomail.v2"
)

var (
	mailDialer = gomail.NewDialer(
		config.SmtpHost,
		config.SmtpPort, 
		config.SmtpUser, 
		config.SmtpPassword,
	)
)