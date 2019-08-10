package sms

import (
	"log"
	"testing"
)

var err error
var client *Client

func init() {
	client, err = New(
		"您的ACCESS KEY",
		"您的SECRET KEY",
		SignName("您的短信签名"),
	)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestClient_Send(t *testing.T) {
	cli := client.Clone(Template("您的短信模板"))
	if err = cli.SendCode("您要发送的手机号码", "12345"); err != nil {
		t.Error(err)
	}
}
