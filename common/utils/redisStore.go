/**
 * redis存取验证码
 */

package utils

import (
	"context"
	"gin-api/config/constant"
	"gin-api/pkg/redis"
	"log"
	"time"
)

var ctx = context.Background()

type RedisStore struct{}

// 存取验证码
func (r RedisStore) Set(id string, value string) {
	key := constant.LOGIN_CODE + id
	err := redis.RedisDb.Set(ctx, key, value, time.Minute*5).Err()
	if err != nil {
		log.Panicln(err.Error())
	}
}

// 获取验证码
func (r RedisStore) Get(id string, clear bool) string {
	key := constant.LOGIN_CODE + id
	val, err := redis.RedisDb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return val
}

// 校验验证码
func (r RedisStore) Verify(id string, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
