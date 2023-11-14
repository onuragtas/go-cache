package redis

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClusterClient(hosts []string) *Client {
	return &Client{Client: redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: hosts,
	})}
}
