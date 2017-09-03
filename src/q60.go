package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/garyburd/redigo/redis"
)

type Artist struct {
	Name string `json:"name"`
	Tags []struct {
		Count int    `json:"count"`
		Value string `json:"value"`
	} `json:"tags"`
	Rating struct {
		Count int `json:"count"`
		Value int `json:"value"`
	} `json:"rating"`
	SortName string `json:"sort_name"`
	Ended    bool   `json:"ended"`
	Gid      string `json:"gid"`
	ID       int    `json:"id"`
	Area     string `json:"area"`
	Aliases  []struct {
		Name     string `json:"name"`
		SortName string `json:"sort_name"`
	} `json:"aliases"`
	Begin struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Date  int `json:"date"`
	} `json:"begin"`
	End struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Date  int `json:"date"`
	} `json:"end"`
	Gender string `json:"gender"`
	Type   string `json:"type"`
}

func main() {
	f, err := os.Open("../data/artist.json")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReaderSize(f, 16384)

	// connect to redis
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		a := &Artist{}

		err = json.Unmarshal(l, a)
		if err != nil {
			panic(err)
		}

		_, err = c.Do("RPUSH", a.Name, a.Area)
		if err != nil {
			panic(err)
		}

	}
}
