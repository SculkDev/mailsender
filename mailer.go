package main

import (
	"crypto/tls"
	"strconv"

	"gopkg.in/gomail.v2"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func SendMail(to, from, subject, htmlBody, textBody string, cfg SMTPConfig) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", textBody)
	m.AddAlternative("text/html", htmlBody)

	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
	d.TLSConfig = &tls.Config{ServerName: cfg.Host}

	return d.DialAndSend(m)
}

func SMTPConfigFromEnv() SMTPConfig {
	port, _ := strconv.Atoi(getEnv("SMTP_PORT", "587"))
	return SMTPConfig{
		Host:     getEnv("SMTP_HOST", "localhost"),
		Port:     port,
		Username: getEnv("SMTP_USERNAME", ""),
		Password: getEnv("SMTP_PASSWORD", ""),
		From:     getEnv("SMTP_FROM", "noreply@example.com"),
	}
}
