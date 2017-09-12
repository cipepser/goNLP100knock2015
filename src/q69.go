package main

import (
	"net/http"
	"sort"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
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

type myForm struct {
	Name      string `form:"name"`
	AliasName string `form:"aliasName"`
	Tags      string `form:"tags"`
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

	// user input
	var name, aliasName string
	var tags []string

	router := gin.Default()
	router.LoadHTMLGlob("../template/*.tmpl")

	// index page
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	// result page
	router.POST("/result", func(c *gin.Context) {
		var fakeForm myForm
		c.Bind(&fakeForm)

		_tags := strings.Split(fakeForm.Tags, " ")

		name = fakeForm.Name
		aliasName = fakeForm.AliasName
		tags = _tags

		// make a query from the user input.
		q := []bson.M{}

		if name != "" {
			q = append(q, bson.M{"name": name})
		}
		if aliasName != "" {
			q = append(q, bson.M{"aliases.name": aliasName})
		}

		if len(tags) > 1 {
			for _, tag := range tags {
				q = append(q, bson.M{"tags.value": tag})
			}
		}
		query := col.Find(bson.M{"$and": q})

		var artists []Artist
		query.All(&artists)
		sort.Slice(artists, func(i, j int) bool {
			return artists[i].Rating.Count < artists[j].Rating.Count
		})

		// send the result to the template.
		c.HTML(http.StatusOK, "result.tmpl", gin.H{
			"artists": artists,
		})
	})

	router.Run(":8080")

}
