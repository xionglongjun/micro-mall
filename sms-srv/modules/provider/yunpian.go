package provider

import (
	"strconv"

	"github.com/shesuyo/yunpian"
)

// YunPian ...
type YunPian struct {
	Name   string
	APIKey string
	Debug  bool
}

// String ...
func (m *YunPian) String() string {
	if m.Name == "" {
		m.Name = "yunpian"
	}
	return m.Name
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
	if !m.Debug {
		result, err = api.SmsSend(smsInfo)
	} else {
		result.Sid = 0
	}
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(result.Sid, 10), nil
}
