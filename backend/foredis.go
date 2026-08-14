package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
	"os"
)

var dbLink redis.Conn
var usingRedis = false

// connectRedis establishes the redis connection and loads existing fortunes.
// Called from main() in a goroutine so it never blocks the HTTP server.
func connectRedis() {
	// Check if REDIS_DNS environment variable is set
	if os.Getenv("REDIS_DNS") == "" {
		fmt.Println("redis config not set")
		return
	}
	var err error
	for i := 0; i < 5; i++ {
		dbLink, err = redis.Dial(
			"tcp",
			fmt.Sprintf("%s:6379", getEnv("REDIS_DNS", "localhost")),
			redis.DialConnectTimeout(2*time.Second),
		)
		if err == nil {
			usingRedis = true
			break
		}
		log.Printf("Attempt %d: redis connection failed: %s", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if !usingRedis {
		log.Println("Failed to connect to redis; serving in-memory fortunes")
		return
	}

	resKeys, err := redis.Values(dbLink.Do("hkeys", "fortunes"))
	if err != nil {
		fmt.Println("redis hkeys failed", err.Error())
		return
	}

	fmt.Printf("*** loading redis fortunes:\n")
	for _, key := range resKeys {
		val, err := dbLink.Do("hget", "fortunes", key)
		if err != nil {
			fmt.Println("redis hget failed", err.Error())
			continue
		}
		idx := fmt.Sprintf("%s", key.([]byte))
		msg := fmt.Sprintf("%s", val.([]byte))
		datastoreDefault.Lock()
		datastoreDefault.m[idx] = fortune{ID: idx, Message: msg}
		datastoreDefault.Unlock()
		fmt.Printf("%s => %s\n", key, val)
	}
}
