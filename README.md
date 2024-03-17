# SSL Reminder

## Description
SSL Reminder takes hosts and mail recipients as inputs and checks the SSL certificate expiration on a scheduled basis using the provided cron expression. Mail recipients will be notified via email when SSL certificate for each host is closing to its expiration date.

## Configurations
- Supports reading environment variable using a `.env` file or through command line

- Sample configurations:

```
# Configurations for email sending
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587                 
SMTP_USER=veron.b@hsenidoutsourcing.com
SMTP_PASSWORD=*********

# Checks the SSL certificate expiration for these hosts/domains
HOST_ADDRESSES=logsilu.knsp.jp,logsiru-dev.practechs.com

# Sends email reminders about expiration to these recipients
MAIL_RECEIVERS=veron.b@hsenidoutsourcing.com,harshana.k@hsenidoutsourcing.com         

# Checks the SSL certificate expiration for each host every 12 hours
SSL_EXPIRE_CHECKER_CRON=0 */12 * * *

# Sends a reminder to renew the certificate 7 days before expiration
REMINDER_DAYS_BEFORE_EXPIRATION=7    
```

## Building the Binary
- Execute command: `make`

## Running the application

### Go Runtime
- Requires Go v1.22 or later
- Execute command: `go run main.go`

### Built Binary
- Execute command: `./dist/ssl-reminder`