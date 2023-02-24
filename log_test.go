package zaplog

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr"`
}

func TestFromContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, Keys[0], uuid.NewString())
	ctx = context.WithValue(ctx, Keys[1], uuid.NewString())

	logger := FromContext(ctx)

	logger.Errorf("this is a test error msg error: %+v", errors.New("name filed is required"))
	logger.Infof("this is a test info msg name: %+v", "saboran")
	logger.Debugf("this is a test debug msg user: %+v", user{
		Name: "saboran",
		Age:  18,
		Addr: "The People's Republic of China",
	})

}
