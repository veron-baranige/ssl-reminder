package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/robfig/cron/v3"
	"github.com/veron-baranige/ssl-reminder/config"
	"github.com/veron-baranige/ssl-reminder/service"
)

func displayLogo() {
	logo := "SSL REMINDER v1.0\nMade by Veron Baranige\n"
	fmt.Println(logo)
}

func main() {
	displayLogo()
	log.SetLevel(log.DebugLevel)
	
	log.Info("Loading environment variables")
	config.LoadEnv()
	
	log.Info("Setting up CRON jobs")
	c := cron.New()

	_, err := c.AddFunc(config.SslCheckerCron, func() {
		log.Info("Running SSL Expire Checker CRON", "time", time.Now())
		for _, host := range config.HostAddresses {
            service.CheckSslCertificateExpiration(host)
        }
	})

	if err != nil {
		log.Error("Failed to add SSL Expire Checker CRON", "err", err)
		return
	}

	c.Start()
	select {}
}