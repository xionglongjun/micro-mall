package handler

import (
	"context"
	"encoding/json"

	"github.com/micro/go-micro/util/log"

	sms "github.com/xionglongjun/micro-mall/sms-api/proto/sms"

	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/xionglongjun/micro-mall/sms-api/client"
)

type Sms struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Sms.Call is called by the API as /sms/call with post body {"name": "foo"}
func (e *Sms) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Sms.Call request")

	// extract the client from the context
	smsClient, ok := client.SmsFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.sms.sms.call", "sms client not found")
	}

	// make request
	response, err := smsClient.Call(ctx, &sms.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.sms.sms.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
