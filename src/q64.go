package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

	// connect to mongodb
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	db := session.DB("MusicBrainz")
	col := db.C("artist")

	prog := 0
	// create databse
	for {
		// read json and decode
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		a := &Artist{}

		err = bson.UnmarshalJSON(l, &a)
		if err != nil {
			panic(err)
		}

		err = col.Insert(a)
		if err != nil {
			panic(err)
		}

		prog++
		if prog%5000 == 0 {
			fmt.Println(prog, "/", 921337)
		}

	}

	a := &Artist{}
	query := col.Find(bson.M{"name": "supercell"})

	t1 := time.Now()
	err = query.One(&a)
	if err != nil {
		panic(err)
	}
	t2 := time.Now()
	fmt.Println("before: ", t2.Sub(t1))

	keys := []string{"name", "aliases.name", "tags.value", "rating.value"}
	for _, k := range keys {
		err = col.EnsureIndexKey(k)
		if err != nil {
			panic(err)
		}
	}

	t1 = time.Now()
	err = query.One(&a)
	if err != nil {
		panic(err)
	}
	t2 = time.Now()
	fmt.Println("after: ", t2.Sub(t1))

	// remove indeces.
	// for _, k := range keys {
	// 	err = col.DropIndex(k)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

}
