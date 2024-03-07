# SSL Reminder

## Description
SSL Reminder takes hosts and mail recipients as inputs and checks the SSL certificate expiration using the provided cron expression. Mail recipients will be notified of the certificate expiration by comparing the days until expiration with the `REMINDER_DAYS_BEFORE_EXPIRATION`.

## Configurations

- Supports environment variable loading either from a `.env` file in the same directory as the binary file or from the directly passed environment variables in command line

- Sample configurations:

```
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=veron.b@hsenidoutsourcing.com
SMTP_PASSWORD=*********

HOST_ADDRESSES=logsilu.knsp.jp,logsiru-dev.practechs.com
MAIL_RECEIVERS=veron.b@hsenidoutsourcing.com,harshana.k@hsenidoutsourcing.com

SSL_EXPIRE_CHECKER_CRON=0 */12 * * *
REMINDER_DAYS_BEFORE_EXPIRATION=7
```
