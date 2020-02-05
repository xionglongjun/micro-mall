package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	send "sms-srv/proto/send"
)

type Send struct{}

func (e *Send) Handle(ctx context.Context, msg *send.CodeRequest) error {
	log.Log("Handler Received message: ", msg.Mobile)
	return nil
}
