package sms

import (
	"fmt"
)

type comResponse struct {
	Code      string
	Message   string
	RequestId string
}

func (cr *comResponse) GetError() error {
	if cr.Code == "OK" {
		return nil
	}
	return fmt.Errorf("alisms: code = %s , message = %s", cr.Code, cr.Message)
}

type sendResponse struct {
	comResponse
	BizId string
}
