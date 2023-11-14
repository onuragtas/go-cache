package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Client struct {
	Client IClient
}

func (c Client) Get(ctx context.Context, key string) *redis.StringCmd {
	return c.Client.Get(ctx, key)
}

func (c Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.Client.Set(ctx, key, value, expiration)
}

func (c Client) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return c.Client.Del(ctx, keys...)
}

func (c Client) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return c.Client.HSet(ctx, key, values...)
}

func (c Client) HGet(ctx context.Context, key string, field string) *redis.StringCmd {
	return c.Client.HGet(ctx, key, field)
}

func (c Client) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	return c.Client.HDel(ctx, key, fields...)
}

func (c Client) MultiGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	return c.Client.HMGet(ctx, key, fields...)
}

func (c Client) MultiSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd {
	return c.Client.HMSet(ctx, key, values)
}

type IClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, key ...string) *redis.IntCmd
	HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	HGet(ctx context.Context, key string, field string) *redis.StringCmd
	HDel(ctx context.Context, key string, field ...string) *redis.IntCmd
	HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd
	HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd
}
