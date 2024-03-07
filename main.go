package main

import (
	"os"

	"github.com/veron-baranige/ssl-reminder/config"
	"go.uber.org/zap"
)

var logger *zap.Logger

func initLogger() {
	os.MkdirAll(config.LogDirectory, os.ModePerm)
	logger, _ = config.GetLoggerConfig().Build()
	defer logger.Sync()
}

func main() {
	initLogger()
}