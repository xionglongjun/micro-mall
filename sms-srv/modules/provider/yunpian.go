package provider

import (
	"strconv"

	"github.com/shesuyo/yunpian"
)

// YunPian ...
type YunPian struct {
	Name   string
	APIKey string
}

// Send ...
func (m *YunPian) Send(mobile, content string) (string, error) {
	api := yunpian.NewYunpianAPI(m.APIKey)
	var (
		smsInfo yunpian.SMSSendInfo
		result  yunpian.SMSResult
		err     error
	)
	smsInfo.Mobile = mobile
	smsInfo.Text = content
	result, err = api.SmsSend(smsInfo)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(result.Sid, 10), nil
}
