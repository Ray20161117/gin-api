/**
* Redis 数据库连接
 */

package redis

import (
	"context"
	config "gin-api/config/yaml_config"

	"github.com/go-redis/redis/v8"
)

var RedisDb *redis.Client

func InitRedis() error {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.Host + ":" + config.Cfg.Redis.Port,
		Password: config.Cfg.Redis.Password,
		DB:       config.Cfg.Redis.Database,
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
