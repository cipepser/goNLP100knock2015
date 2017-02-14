package main

import (
	"encoding/json"
	"fmt"
	"os"
	"bufio"
	"io"
	"regexp"
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
	
	var txt string
	for _, article := range articles {
		if article.Title == "イギリス" {
			txt = article.Text
		}
	}
	
	reg, _ := regexp.Compile(`(?m)^=+.*=+`)
	for _, v := range reg.FindAll([]byte(txt), -1) {
		for i := 0; i < len(v); i++ {
			if string(v[i]) != "=" {
				fmt.Println(string(v[i:len(v) - i]), i - 1)
				break
			}
		}		
	}	
}