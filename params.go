package sms

import (
	"encoding/json"
	"errors"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type BatchItem struct {
	Sign   string
	Mobile string
	Params map[string]interface{}
}

type BatchItems []BatchItem

func (items BatchItems) applyTo(req *requests.CommonRequest) error {
	if len(items) == 0 {
		return errors.New("alisms: the batch send items can not be empty")
	}
	signs := make([]string, len(items))
	mobiles := make([]string, len(items))
	params := make([]map[string]interface{}, len(items))
	for i, item := range items {
		if item.Mobile == "" {
			continue
		} else if item.Sign == "" {
			item.Sign = req.QueryParams["SignName"]
		}
		signs[i] = item.Sign
		mobiles[i] = item.Mobile
		params[i] = item.Params
	}
	if sbs, err := json.Marshal(signs); err != nil {
		return err
	} else if mbs, err := json.Marshal(mobiles); err != nil {
		return err
	} else if pns, err := json.Marshal(params); err != nil {
		return err
	} else {
		delete(req.QueryParams, "SignName")
		req.QueryParams["SignNameJson"] = string(sbs)
		req.QueryParams["PhoneNumberJson"] = string(mbs)
		req.QueryParams["TemplateParamJson"] = string(pns)
	}
	return nil
}
