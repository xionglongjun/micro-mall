package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	user "user-srv/proto/user"
)

type User struct{}

func (e *User) Handle(ctx context.Context, msg *user.CreateRequest) error {
	log.Log("Handler Received message: ", msg.Name)
	return nil
}

func Handler(ctx context.Context, msg *user.CreateRequest) error {
	log.Log("Function Received message: ", msg.Name)
	return nil
}
