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

	rep, err := redis.Strings(c.Do("LRANGE", os.Args[1], "0", "-1"))
	if err != nil || len(rep) == 0 {
		fmt.Println("`", os.Args[1], "` is not found.")
	} else {
		for _, r := range rep {
			fmt.Println(r)
		}
	}
}
