package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("SET", "Shirdon", "18")
	if err != nil {
		fmt.Println("redis set error:", err)
	}

	name, err := redis.String(conn.Do("GET", "Shirdon"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("Get name: %s \n", name)
	}
}
