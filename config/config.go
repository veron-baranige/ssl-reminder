package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	SmtpHost              string
	SmtpUser              string
	SmtpPassword          string
	SmtpPort              int
	HostAddresses         []string
	ReceiverMailAddresses []string
	DaysToRemindFrom      int
	SslCheckerCron        string
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	SmtpPort, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	SmtpPassword = os.Getenv("SMTP_PASSWORD")
	SmtpUser = os.Getenv("SMTP_USER")
	SmtpHost = os.Getenv("SMTP_HOST")

	ReceiverMailAddresses = strings.Split(os.Getenv("MAIL_RECEIVERS"), ",")
	HostAddresses = strings.Split(os.Getenv("HOST_ADDRESSES"), ",")
	DaysToRemindFrom, _ = strconv.Atoi(os.Getenv("REMINDER_DAYS_BEFORE_EXPIRATION"))
	SslCheckerCron = os.Getenv("SSL_EXPIRE_CHECKER_CRON")

	return nil
}
