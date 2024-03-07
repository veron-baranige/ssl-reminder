package service

import (
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/veron-baranige/ssl-reminder/config"
	"github.com/veron-baranige/ssl-reminder/mail"
)

var (
	debugLogger = log.NewWithOptions(os.Stderr, log.Options{
		Level:           log.DebugLevel,
		TimeFormat:      time.DateTime,
		ReportTimestamp: true,
	})
)

func CheckSslCertificateExpiration(host string) {
	log.Info("Dialing host", "host", host)

	parsedHostAddr := fmt.Sprintf("%s:%d", host, 443)
	conn, err := tls.Dial("tcp", parsedHostAddr, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Error("Failed to retrieve TLS connection", "host", host, "err", err)
		return
	}

	cert := conn.ConnectionState().PeerCertificates[0]
	debugLogger.Debug("Certificate expiration date:", "cert.NotAfter", cert.NotAfter)

	timeUntilExp := time.Until(cert.NotAfter)
	daysUntilExp := int(timeUntilExp.Hours() / 24)
	debugLogger.Debug("Days until ceritifcate expiration: ", "daysUntilExp", daysUntilExp)

	if daysUntilExp <= 0 {
		log.Warn("Certificate has expired", "host", host)
		log.Info("Sending reminder mail to renew the certificate")

		if err := mail.SendMail(
			fmt.Sprintf("SSL Reminder - %s", host),
			fmt.Sprintf(
				"SSL certificate has been expired for the host: %s\n" + 
				"Please renew the certificate immediately.", 
				host,
			),
		); err != nil {
			log.Error("Failed to send mail", "err", err)
		}
	} else if daysUntilExp <= config.DaysToRemindFrom {
		log.Info("Sending reminder mail to renew the certificate")
		if err := mail.SendMail(
			fmt.Sprintf("SSL Reminder - %s", host),
			fmt.Sprintf(
				"SSL certificate is about to expire for the host: %s\n"+
				"Days until expiration: %d\n"+
				"Please renew the certificate immediately.", 
				host, 
				daysUntilExp,
			),
		); err != nil {
			log.Error("Failed to send mail", "err", err)
		}
	}
}
