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
	// user input
	name := "Queen"
	aliasName := "Queen"
	// tags := []string{"hard_rock"}

	// connect to mongodb
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	db := session.DB("MusicBrainz")
	col := db.C("artist")

	// q := `"name":`
	// query := col.Find(bson.M{"name": name})

	// q := `"name": name, "aliases.name": aliasName, "tags.value": "related-akb48", "kamen rider w"`
	q := bson.M{
		"name":         name,
		"aliases.name": aliasName,
	}
	// $and: [
	// 	{"tags.value": "kamen rider w"},
	// 	{"tags.value": "related-akb48"}
	// ]

	query := col.Find(q)

	var artists []Artist
	query.All(&artists)

	for _, a := range artists {
		// fmt.Println(a.Tags[0].Value)
		fmt.Println(a)
	}

	// sort.Slice(artists, func(i, j int) bool {
	// 	return artists[i].Rating.Count < artists[j].Rating.Count
	// })
	//
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(artists[len(artists)-i-1].Name)
	// }
}
