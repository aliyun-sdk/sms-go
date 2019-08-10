package sms

import (
	"flag"
	"log"
	"testing"
)

var err error
var client *Client
var ak, sk, sn, tc string

func init() {
	flag.StringVar(&ak, "access_key", "", "您的ACCESS KEY")
	flag.StringVar(&sk, "secret_key", "", "您的SECRET KEY")
	flag.StringVar(&sn, "sign_name", "", "您的短信发送签名")
	flag.StringVar(&tc, "template_code", "", "您的短信发送模板")
	flag.Parse()
	client, err = New(ak, sk, SignName(sn), Template(tc))
	if err != nil {
		log.Fatalln(err)
	}
}

func TestClient_Send(t *testing.T) {
	err = client.SendCode("17757171483", "12345")
	if err != nil {
		t.Error(err)
	}
}

func TestClient_SendBatch(t *testing.T) {
	items := []BatchItem{
		{
			// "您的短信签名, 若已全局配置则可留空"
			Sign:   "",
			Mobile: "17757171482",
			Params: map[string]interface{}{"code": "01234"},
		},
		{
			// "您的短信签名, 若已全局配置则可留空"
			Sign:   "",
			Mobile: "17757171483",
			Params: map[string]interface{}{"code": "56789"},
		},
	}
	err = client.SendBatch(items)
	if err != nil {
		t.Error(err)
	}
}
