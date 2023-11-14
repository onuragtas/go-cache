package redis

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(host string) *Client {

	return &Client{Client: redis.NewClient(&redis.Options{
		Addr: host,
	})}
}
