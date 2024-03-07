package config

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	SmtpHost = "smtp.gmail.com"
	SmtpPort = 587
	LogDirectory = "./logs"
	InfoLogPath = "./logs/info.log"
	ErrorLogPath = "./logs/error.log"
)

var (
	SmtpUser      = os.Getenv("SMTP_USER")
	SmtpPassword  = os.Getenv("SMTP_PASSWORD")
	MailReceivers = os.Getenv("MAIL_RECEIVERS")
	config        = zap.NewDevelopmentConfig()
)

func GetLoggerConfig() zap.Config {
	config.EncoderConfig.TimeKey = "timestamp"
    config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	config.DisableCaller = true
    config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.OutputPaths = []string{"stderr", InfoLogPath}
	return config
}
