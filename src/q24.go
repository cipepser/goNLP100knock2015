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
	
	reg, _ := regexp.Compile(`(File|ファイル):.*?\|`)
	for _, v := range reg.FindAll([]byte(txt), -1) {
		if string(v[0]) == "F" {
			fmt.Println(string(v[len("File:"): len(v) - 1]))
		} else {
			fmt.Println(string(v[len("ファイル:"): len(v) - 1]))
		}
	}
	
}