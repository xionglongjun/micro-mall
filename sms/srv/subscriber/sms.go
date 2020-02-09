package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	sms "github.com/xionglongjun/micro-mall/sms/srv/proto/sms"
)

type Sms struct{}

func (e *Sms) Handle(ctx context.Context, msg *sms.ListRequest) error {
	log.Log("Handler Received message: ", msg.Page)
	return nil
}
