package mail

import (
	"encoding/base64"
	"net/mail"
	"fmt"
	"net/smtp"
)

type SMTP struct {
	Host     string
	Password string
	Email    string
	Port     int
}

func (inst *SMTP) SendMail(to, title, message string) error{
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	fromAddress := mail.Address{Name:"Sender", Address:inst.Email}
	toAddress := mail.Address{Name:"To", Address:to}
	header := map[string]string{
		"From": fromAddress.String(),
		"To": toAddress.String(),
		"Subject": fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(title))),
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=UTF-8",
		"Content-Transfer-Encoding": "base64",
	}
	emailContent := ""
	for k, v := range header {
		emailContent += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	emailContent += "\r\n" + b64.EncodeToString([]byte(message))
	auth := smtp.PlainAuth("", inst.Email, inst.Password, inst.Host)
	err := smtp.SendMail( fmt.Sprintf("%s:%d", inst.Host, inst.Port), auth, inst.Email, []string{toAddress.Address},
	[]byte(emailContent))
	return err
}
