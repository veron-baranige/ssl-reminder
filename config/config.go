package config

import "os"

const (
	SmtpHost = "smtp.gmail.com"
	SmtpPort = 587
)

var (
	SmtpUser      = os.Getenv("SMTP_USER")
	SmtpPassword  = os.Getenv("SMTP_PASSWORD")
	MailReceivers = os.Getenv("MAIL_RECEIVERS")
)
