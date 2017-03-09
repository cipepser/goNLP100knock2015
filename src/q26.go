package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Article struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

func main() {
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
	
	// 基本情報だけ抜き出し
	reg := regexp.MustCompile(`{{基礎情報 国[\s\S]*\n}}`)
	txt = string(reg.FindAll([]byte(txt), -1)[0])
	txt = strings.Replace(txt, "{{基礎情報 国\n", "", 1)
	
	// 前処理
	// `\n|`でマッチさせると偶数回目がfindできないため`|`をダブらせる
	txt = strings.Replace(txt, "\n|", "\n|\n|", -1)
	
	// フィールドごとに分ける
	reg = regexp.MustCompile(`(?m)^\|[\s\S]*?\n\|`)

	// Mapに格納
	m := make(map[string]string)
	regEmphasis := regexp.MustCompile(`'{2,5}`)
	for _, v := range reg.FindAll([]byte(txt), -1) {
		s := string(v[1: len(v) - 2])
		strs := strings.Split(s, " = ")

		m[strs[0]] = regEmphasis.ReplaceAllString(strs[1], "")
	}
	
	fmt.Println(m["確立形態4"])
}
