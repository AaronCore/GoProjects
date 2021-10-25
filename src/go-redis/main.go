package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func redisInit() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 1000,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func main() {
	err := redisInit()
	if err != nil {
		fmt.Println("connect redis failed,err：", err)
		return
	}

	fmt.Println("1、redis key")
	key1 := "test1"
	val, _ := rdb.Exists(ctx, key1).Result()
	fmt.Println(val)
	val1, _ := rdb.Del(ctx, key1).Result()
	fmt.Println(val1)

	fmt.Println("2、redis set和get")
	key2 := "test2"
	err = rdb.Set(ctx, key2, "aaaa", time.Second*10).Err()
	if err != nil {
		fmt.Println("set failed,err：", err)
		panic(err)
	}
	v, err := rdb.Get(ctx, key2).Result()
	if err != nil {
		fmt.Println("get failed,err：", err)
		panic(err)
	}
	fmt.Println(v)

	fmt.Println("3、redis ZAdd")
	key3 := "test3"
	items := []*redis.Z{
		&redis.Z{Score: 10, Member: "10"},
		&redis.Z{Score: 11, Member: "11"},
		&redis.Z{Score: 17, Member: "17"},
		&redis.Z{Score: 13, Member: "13"},
	}
	num, err := rdb.ZAdd(ctx, key3, items...).Result()
	if err != nil {
		fmt.Println("ZAdd failed,err：", err)
		panic(err)
	}
	fmt.Println(num)

	newScore, err := rdb.ZIncrBy(ctx, key3, 2, "10").Result()
	if err != nil {
		fmt.Println("ZIncrBy failed,err：", err)
		panic(err)
	}
	fmt.Println(newScore)
}
