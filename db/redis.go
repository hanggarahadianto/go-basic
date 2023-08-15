package db

import (
	"fmt"
	"go-basic/utils"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func RedisInit()error{

	config, _:= utils.LoadConfig(".")
	db,err := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
		Password : "",
		DB: 2,
	})
	 

	fmt.Println("connect to redist")
	rdb = db

	return nil
}

func RedisConnect() *redis.Client{
	return rdb
}