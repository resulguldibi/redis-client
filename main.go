package main

import (
	"fmt"
	"resulguldibi/redis-client/factory"
)

func main() {

	redisClientFactory := factory.NewRedisClientFactory("localhost:6379", "")
	redisClient := redisClientFactory.GetRedisClient()

	data := make(map[string]interface{})
	data["key1"] = "a"
	data["key2"] = "b"
	res, err := redisClient.HMSet("resul", data)

	fmt.Println(res)
	fmt.Println(err)

	res2, err2 := redisClient.HMGet("resul", "key1")

	fmt.Println(res2)
	fmt.Println(err2)
}
