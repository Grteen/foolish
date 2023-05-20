package rdb

import (
	"be/pkg/constants"
	"be/pkg/kafka"
	"context"
	"time"
)

// 存储注册验证码
func SetRegisterVerify(ctx context.Context, email, verify string) error {
	if err := RDB.Set(ctx, constants.RdbRegisterVerify+email, verify, time.Duration(constants.RdbRegisterVerifyExpiration*constants.ChangeToRedis)).Err(); err != nil {
		kafka.ErrorLog(err.Error())
		return err
	}
	return nil
}

// 查询注册验证码
func QueryRegisterVerify(ctx context.Context, email string) (string, error) {
	res := RDB.Get(ctx, constants.RdbRegisterVerify+email)
	return res.Result()
}
