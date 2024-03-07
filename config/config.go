package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
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

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		if err.Error() == "open .env: no such file or directory" {
			log.Warn("Failed to load environment variables using .env file")
		} else {
			log.Error("Failed to load environment variables", "err", err)
			os.Exit(1)
		}
	}

	SmtpPort, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	SmtpPassword = os.Getenv("SMTP_PASSWORD")
	SmtpUser = os.Getenv("SMTP_USER")
	SmtpHost = os.Getenv("SMTP_HOST")

	ReceiverMailAddresses = strings.Split(os.Getenv("MAIL_RECEIVERS"), ",")
	HostAddresses = strings.Split(os.Getenv("HOST_ADDRESSES"), ",")
	DaysToRemindFrom, _ = strconv.Atoi(os.Getenv("REMINDER_DAYS_BEFORE_EXPIRATION"))
	SslCheckerCron = os.Getenv("SSL_EXPIRE_CHECKER_CRON")
}
