package sms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type Options []Option

func (opts Options) applyTo(req *requests.CommonRequest) {
	for _, opt := range opts {
		opt(req)
	}
}

type Option func(req *requests.CommonRequest)

func Scheme(s string) Option {
	return func(req *requests.CommonRequest) {
		req.Scheme = s
	}
}

func Domain(d string) Option {
	return func(req *requests.CommonRequest) {
		req.Domain = d
	}
}

func Version(v string) Option {
	return func(req *requests.CommonRequest) {
		req.Version = v
	}
}

func Method(m string) Option {
	return func(req *requests.CommonRequest) {
		req.Method = m
	}
}

func Action(a string) Option {
	return func(req *requests.CommonRequest) {
		req.ApiName = a
	}
}

func QueryParam(k, v string) Option {
	return func(req *requests.CommonRequest) {
		req.QueryParams[k] = v
	}
}

func Mobile(ss string) Option {
	return QueryParam("PhoneNumbers", ss)
}

func Parameter(v map[string]string) Option {
	param, _ := json.Marshal(v)
	return QueryParam("TemplateParam", string(param))
}

func SignName(n string) Option {
	return QueryParam("SignName", n)
}

func Template(t string) Option {
	return QueryParam("TemplateCode", t)
}
