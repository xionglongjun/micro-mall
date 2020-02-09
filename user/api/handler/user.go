package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/util/log"

	"user-api/client"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/go-micro/api/proto"
	user "path/to/service/proto/user"
)

type User struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// User.Call is called by the API as /user/call with post body {"name": "foo"}
func (e *User) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received User.Call request")

	// extract the client from the context
	userClient, ok := client.UserFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.user.user.call", "user client not found")
	}

	// make request
	response, err := userClient.Call(ctx, &user.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.user.user.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
