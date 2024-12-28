/**
* Redis 数据库连接
 */

package redis

import (
	"context"
	"gin-api/config/yaml"

	"github.com/go-redis/redis/v8"
)

var RedisDb *redis.Client

func InitRedis() error {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     yaml.Cfg.Redis.Host + ":" + yaml.Cfg.Redis.Port,
		Password: yaml.Cfg.Redis.Password,
		DB:       yaml.Cfg.Redis.Database,
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
