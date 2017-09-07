package main

import (
	"fmt"

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
	// connect to mongodb
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	db := session.DB("MusicBrainz")
	col := db.C("artist")

	query := col.Find(bson.M{"name": "Queen"})

	var artists []Artist
	query.All(&artists)

	for _, a := range artists {
		fmt.Println(a)
	}

}
