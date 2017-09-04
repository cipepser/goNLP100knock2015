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

	// get keys
	keys, err := redis.Strings(c.Do("KEYS", "*"))
	if err != nil {
		panic(err)
	}

	// count up
	cnt := 0
	for _, k := range keys {
		rep, err := redis.Strings(c.Do("LRANGE", k, "0", "-1"))
		if err != nil {
			panic(err)
		} else {
			for _, r := range rep {
				if r == "Japan" {
					cnt++
				}
			}
		}
	}

	fmt.Println(cnt)
}
