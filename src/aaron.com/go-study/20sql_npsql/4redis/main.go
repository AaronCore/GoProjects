package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var redisDb *redis.Client
var ctx = context.Background()

func redisInit() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 1000,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = redisDb.Ping(ctx).Result()
	return err
}

func main() {
	err := redisInit()
	if err != nil {
		fmt.Println("connect redis failed,err：", err)
		return
	}
	fmt.Println("1、set、get")
	err = redisDb.Set(ctx, "k1", 100, time.Second*10).Err()
	if err != nil {
		fmt.Println("set failed,err：", err)
		panic(err)
	}
	v, err := redisDb.Get(ctx, "k1").Result()
	if err != nil {
		fmt.Println("get failed,err：", err)
		panic(err)
	}
	fmt.Println(v)

	fmt.Println("2、zset、zget")
	//items := []*redis.Z{
	//	&redis.Z{Score: 10, Member: "10"},
	//	&redis.Z{Score: 11, Member: "11"},
	//	&redis.Z{Score: 17, Member: "17"},
	//	&redis.Z{Score: 13, Member: "13"},
	//}
	//num, err := redisDb.ZAdd(ctx, "ZSET1", items...).Result()
	//if err != nil {
	//	fmt.Println("zadd failed,err：", err)
	//	panic(err)
	//}
	//fmt.Println(num)

	newScore, err := redisDb.ZIncrBy(ctx, "ZSET1", 2, "10").Result()
	if err != nil {
		fmt.Println("ZIncrBy failed,err：", err)
		panic(err)
	}
	fmt.Println(newScore)
}
