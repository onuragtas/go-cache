package cache

import (
	"context"
	"errors"
	"time"

	"github.com/onuragtas/go-cache/cache/serializer"
	"github.com/onuragtas/go-cache/database/redis"
)

func (c *Adapter) Get(serializer serializer.ISerializer, key string) error {
	data, err := c.client.Get(context.TODO(), key).Result()
	if err != nil {
		return err
	}

	err = serializer.Deserialize(data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Adapter) Set(serializer serializer.ISerializer, key string, value interface{}, ttl time.Duration) error {
	data := serializer.Serialize(value)

	if data == "" {
		return errors.New("data is empty")
	}

	return c.client.Set(context.TODO(), key, data, ttl).Err()
}

func (c *Adapter) Del(keys ...string) error {
	return c.client.Del(context.TODO(), keys...).Err()
}

func (c *Adapter) HSet(serializer serializer.ISerializer, key string, value ...interface{}) error {
	data := serializer.Serialize(value[1])

	if data == "" {
		return errors.New("data is empty")
	}

	return c.client.HSet(context.TODO(), key, value[0], data).Err()
}

func (c *Adapter) HGet(serializer serializer.ISerializer, key string, field string) error {
	data, err := c.client.HGet(context.TODO(), key, field).Result()
	if err != nil {
		return err
	}

	err = serializer.Deserialize(data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Adapter) HDel(key string, fields ...string) error {
	return c.client.HDel(context.TODO(), key, fields...).Err()
}

func (c *Adapter) MultiGet(serializer *serializer.Serializer, key string, fields ...string) (interface{}, error) {
	unserialized := false
	var err error
	data, err := c.client.MultiGet(context.TODO(), key, fields...).Result()
	if err != nil {
		return nil, err
	}

	for _, item := range data {
		if item != nil {
			err = serializer.DeserializeType(item.(string))
			if err != nil {
				return nil, err
			}
			unserialized = true
		}
	}

	if unserialized == false {
		return nil, errors.New("data is empty")
	}

	return serializer.GetOut(), err
}

func (c *Adapter) MultiSet(serializer *serializer.Serializer, key string, values ...interface{}) error {
	// serialize []map[string]interface{} to []interface{}
	for i := 0; i < len(values); i++ {
		if i%2 == 0 {
			continue
		}
		values[i] = serializer.Serialize(values[i])
	}

	return c.client.MultiSet(context.TODO(), key, values...).Err()
}

func (c *Adapter) DeleteHashWithPattern(key, pattern string, offset int64, count int64) error {
	var err error
	if pattern == "" {
		err = c.deleteWithScan(key, offset, count)
	} else {
		err = c.deleteWithHash(key, pattern, offset, count)
	}
	return err
}

func (c *Adapter) deleteWithScan(key string, offset int64, count int64) error {
	var keys []string
	cursor := uint64(offset)

	for {
		list, nextCursor, err := c.client.Scan(context.TODO(), cursor, key, count)
		if err != nil {
			return err
		}

		if len(list) > 0 {
			keys = append(keys, list...)
		}

		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}

	if len(keys) != 0 {
		return c.client.Del(context.TODO(), keys...).Err()
	}

	return nil
}

func (c *Adapter) deleteWithHash(key, pattern string, offset int64, count int64) error {
	var keys []string
	cursor := uint64(offset)

	for {
		list, nextCursor, err := c.client.HScan(context.TODO(), key, cursor, pattern, count)
		if err != nil {
			return err
		}

		for i, foundKey := range list {
			if i%2 != 0 {
				continue
			}
			keys = append(keys, foundKey)
		}

		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}
	if len(keys) != 0 {
		return c.client.HDel(context.TODO(), key, keys...).Err()
	}
	return nil
}

func NewCacheAdapter(hosts []string) *Adapter {
	if len(hosts) > 1 {
		return &Adapter{
			client: redis.NewRedisClusterClient(hosts),
		}
	}

	return &Adapter{
		client: redis.NewRedisClient(hosts[0]),
	}
}
