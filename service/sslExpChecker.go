package service

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/charmbracelet/log"
	"github.com/veron-baranige/ssl-reminder/config"
	"github.com/veron-baranige/ssl-reminder/mail"
)

const (
	timeout = 20 * time.Second
)

func CheckSslCertificateExpiration(host string) {
	log.Info("Dialing host", "host", host)

	parsedHostAddr := fmt.Sprintf("%s:%d", host, 443)
	conn, err := tls.DialWithDialer(
		&net.Dialer{Timeout: timeout}, 
		"tcp", 
		parsedHostAddr,
		&tls.Config{InsecureSkipVerify: true},
	)
	if err != nil {
		log.Error("Failed to retrieve TLS connection", "host", host, "err", err)
		return
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	log.Debug("Certificate expiration date:", "cert.NotAfter", cert.NotAfter)

	timeUntilExp := time.Until(cert.NotAfter)
	daysUntilExp := int(timeUntilExp.Hours() / 24)
	log.Debug("Days until ceritifcate expiration: ", "daysUntilExp", daysUntilExp)

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
