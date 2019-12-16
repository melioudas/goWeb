package myredis

import (
	"dataPlatform/model"
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
	"strings"
	"time"
)

var RedisClient *redis.ClusterClient
func InitRedis() {
	fmt.Println("===>INIT REDIS!")
	//单机模式
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	//Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	//集群模式
	kAddr := model.ConfigParam["redis-client"]
	addr := strings.Split(kAddr, ",")
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              addr,
		DialTimeout:        10 * time.Second,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolSize:           10,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        time.Minute,
		IdleCheckFrequency: 100 * time.Millisecond,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	RedisClient = rdb
}
