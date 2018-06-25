package factory

import (
	"resulguldibi/redis-client/entity"
)

type IRedisClientFactory interface {
	GetRedisClient() entity.IRedisClient
}

type RedisClientFactory struct {
	address  string
	password string
}

func NewRedisClientFactory(address string, password string) IRedisClientFactory {
	return &RedisClientFactory{address: address, password: password}
}

func (f *RedisClientFactory) GetRedisClient() entity.IRedisClient {
	return entity.NewRedisClient(f.address, f.password)
}
