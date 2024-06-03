package cache

import (
	"time"

	"github.com/onuragtas/go-cache/cache/serializer"
	"github.com/onuragtas/go-cache/database/redis"
)

type Adapter struct {
	IAdapter
	client *redis.Client
}

type IAdapter interface {
	Get(serialize serializer.ISerializer, key string) error
	Set(serialize serializer.ISerializer, key string, value interface{}, ttl time.Duration) error
	Del(keys ...string) error
	HSet(serialize serializer.ISerializer, key string, values ...interface{}) error
	HGet(serialize serializer.ISerializer, key string, field string) error
	MultiGet(serialize *serializer.Serializer, key string, fields ...string) (interface{}, error)
	MultiSet(serialize *serializer.Serializer, key string, values ...interface{}) error
	HDel(key string, field ...string) error
	DeleteHashWithPattern(key, pattern string, offset uint64, count int64) error
}
