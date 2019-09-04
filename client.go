package sms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type Client struct {
	client  *sdk.Client
	options Options
}

func (c *Client) Clone(fns ...Option) *Client {
	return &Client{
		client:  c.client,
		options: append(c.options, fns...),
	}
}

// Send 短信发送
// 支持对多个手机号码发送短信,手机号码之间以英文逗号分隔,上限为1000个手机号码
// 详见 https://help.aliyun.com/document_detail/101414.html
func (c *Client) Send(fns ...Option) error {
	fns = append(fns, Method("POST"))
	fns = append(fns, Action("SendSms"))
	res, err := c.client.ProcessCommonRequest(c.genRequest(fns...))
	if err != nil {
		return err
	}
	var sendRes sendResponse
	if err := json.Unmarshal(res.GetHttpContentBytes(), &sendRes); err != nil {
		return err
	}
	return sendRes.GetError()
}

// SendBatch 批量发送短信
// 与Send不同的是,SendBatch支持不同的发送参数和签名,上限为100个手机号码
// 详见 https://help.aliyun.com/document_detail/102364.html
func (c *Client) SendBatch(items BatchItems) error {
	req := c.genRequest(
		Method("POST"),
		Action("SendBatchSms"),
	)
	if err := items.applyTo(req); err != nil {
		return err
	}
	res, err := c.client.ProcessCommonRequest(req)
	if err != nil {
		return err
	}
	var sendRes sendResponse
	if err := json.Unmarshal(res.GetHttpContentBytes(), &sendRes); err != nil {
		return err
	}
	return sendRes.GetError()
}

func (c *Client) genRequest(fns ...Option) *requests.CommonRequest {
	req := requests.NewCommonRequest()
	c.options.applyTo(req)
	Options(fns).applyTo(req)
	return req
}

func New(ak, sk string, fns ...Option) (*Client, error) {
	cli, err := sdk.NewClientWithAccessKey("default", ak, sk)
	if err != nil {
		return nil, err
	}
	opts := []Option{
		Scheme("http"),
		Domain("dysmsapi.aliyuncs.com"),
		Version("2017-05-25"),
	}
	return &Client{client: cli, options: append(opts, fns...)}, nil
}
