package mail

import (
	"net/smtp"
	"strings"
)

func SendMail(host string, port int, username, password, from string, to []string, subject, body string) error {
	auth := smtp.PlainAuth("", username, password, host)
	msg := []byte("To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=\"utf-8\"\r\n\r\n" +
		body)
	return smtp.SendMail(host+":"+string(rune(port)), auth, from, to, msg)
}
