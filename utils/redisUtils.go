package utils

import (
	"bara-playdate-api/exception"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-redis/redis"
)

func NewRedis(config Config) *redis.Client {
	host := config.RedisHost
	port := config.RedisPort
	maxPoolSize, err := strconv.Atoi(config.RedisMaxSize)
	minIdlePoolSize, err := strconv.Atoi(config.RedisMinIdleSize)
	exception.PanicLogging(err)

	redisStore := redis.NewClient(&redis.Options{
		Addr:         host + ":" + port,
		PoolSize:     maxPoolSize,
		MinIdleConns: minIdlePoolSize,
	})
	return redisStore
}

func SetCache[T any](cacheManager *redis.Client, ctx context.Context, prefix string, key string, executeData func(context.Context, string) (T, error)) *T {
	var data []byte
	var object T
	if err := cacheManager.Get(prefix + "_" + key).Scan(&data); err == nil {
		err := json.Unmarshal(data, &object)
		exception.PanicLogging(err)

		return &object
	}
	value, err := executeData(ctx, key)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	cacheValue, err := json.Marshal(value)
	exception.PanicLogging(err)

	if err := cacheManager.Set(prefix+"_"+key, cacheValue, -1).Err(); err != nil {
		exception.PanicLogging(err)
	}
	return &value
}
