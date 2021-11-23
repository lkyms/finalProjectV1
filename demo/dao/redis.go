package dao

import (
	"demo/util"
	"strconv"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var Rdb *redis.Client

// 初始化连接
func init() {
	db, _ := strconv.Atoi(util.GetConfig("redis.db"))
	Rdb = redis.NewClient(&redis.Options{
		Addr:     util.GetConfig("redis.address"),
		Password: util.GetConfig("redis.pwd"), // no password set
		DB:       db,                          // use default DB
	})

	_, err := Rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
}
