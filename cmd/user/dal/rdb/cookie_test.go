package rdb

import (
	"context"
	"testing"
)

func TestCookie(t *testing.T) {
	RedisInit()
	err := SetCookie(context.Background(), "cookie-test", "a good value", 1000*1000*1000*60)
	if err != nil {
		t.Error(err)
	}

	res, err := QueryCookie(context.Background(), "cookie-test")
	if err != nil {
		t.Error(err)
	}

	if res != "a good value" {
		t.Fail()
	}
}
