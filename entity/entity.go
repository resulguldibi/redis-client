package entity

import (
	"github.com/go-redis/redis"
)

type IRedisClient interface {
	HMGet(key string, fields ...string) (string, error)
	HMSet(key string, fields map[string]interface{}) (string, error)
	HDel(key string, fields ...string) (int64, error)
	HIncrBy(key, field string, incr int64) (int64, error)
	HGet(key, field string) (string, error)
	HSet(key, field string, value interface{}) (bool, error)
}

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(address string, password string) IRedisClient {
	return &RedisClient{client: redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       0,        // use default DB
	})}
}

func (c *RedisClient) HMGet(key string, fields ...string) (string, error) {
	var result string = ""
	response, err := c.client.HMGet(key, fields...).Result()
	if err == nil && response != nil && len(response) > 0 && response[0] != nil {
		result = response[0].(string)
	}

	return result, err
}

func (c *RedisClient) HMSet(key string, fields map[string]interface{}) (string, error) {
	var result string = ""

	response, err := c.client.HMSet(key, fields).Result()

	if err == nil {
		result = response
	}

	return result, err
}

func (c *RedisClient) HDel(key string, fields ...string) (int64, error) {
	var result int64 = 0

	response, err := c.client.HDel(key, fields...).Result()

	if err == nil {
		result = response
	}

	return result, err
}

func (c *RedisClient) HIncrBy(key, field string, incr int64) (int64, error) {
	var result int64 = 0

	response, err := c.client.HIncrBy(key, field, incr).Result()

	if err == nil {
		result = response
	}

	return result, err
}

func (c *RedisClient) HGet(key string, field string) (string, error) {
	var result string = ""
	result, err := c.client.HGet(key, field).Result()
	return result, err
}

func (c *RedisClient) HSet(key string, field string, value interface{}) (bool, error) {
	var result bool = false

	result, err := c.client.HSet(key, field, value).Result()

	return result, err
}
