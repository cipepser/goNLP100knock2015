package main

import (
	"encoding/json"
	"fmt"
	"os"
	"bufio"
	"io"
)

type Article struct {
	Text 	string	`json:"text"`
	Title string	`json:"title"`
}


func main()  {
	articles := []Article{}
	
	f, err := os.Open("../data/jawiki-country.json")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	
	r := bufio.NewReader(f)
	for {
		b, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		a := Article{}
		json.Unmarshal([]byte(b), &a)
		articles = append(articles, a)
	}
	
	for _, article := range articles {
		if article.Title == "イギリス" {
			fmt.Println(article.Text)
		}
	}
	
}