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

	cnt := 0
	for _, k := range keys {
		v, err := redis.Stringa(c.Do("GET", k))
		if err != nil {
			panic(err)
		}
		if v == "Japan" {
			cnt++
		}
	}

	fmt.Println(cnt)
}
