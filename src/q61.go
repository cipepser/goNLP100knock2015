package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
)

func main() {
	// connect to redis
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	if len(os.Args) != 1 {
		errors.New("please input artist name.")
	}

	res, err := redis.String(c.Do("GET", os.Args[1]))
	if err != nil {
		fmt.Println("`", os.Args[1], "` is not found.")
	} else {
		fmt.Println(res)
	}
}
