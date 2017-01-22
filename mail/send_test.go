package mail

import (
	"testing"
	"github.com/bobcyw/goUtils"
)

type EmailTo struct {
	To      string
	Title   string
	Content string
}

func TestSMTP_SendMail(t *testing.T) {
	conf := goUtils.FileString("sensitive/test_send.json")
	var smtp SMTP
	if err := conf.ReadObj(&smtp); err != nil {
		t.Fatal(err.Error())
	}
	sendConf := goUtils.FileString("sensitive/test_to.json")
	var sendContent EmailTo
	if err := sendConf.ReadObj(&sendContent); err != nil{
		t.Fatal(err.Error())
	}
	if err := smtp.SendMail(sendContent.To, sendContent.Title, sendContent.Content); err != nil {
		t.Fatal(err.Error())
	}
}