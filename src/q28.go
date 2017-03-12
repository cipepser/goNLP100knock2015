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

	// 注記のあとには`\n|`がないので追加しておく
	txt = strings.Replace(txt, "\n}}", "\n|}}", -1)

	// フィールドごとに分ける
	reg = regexp.MustCompile(`(?m)^\|[\s\S]*?\n\|`)

	// Mapに格納
	m := make(map[string]string)
	regEmphasis := regexp.MustCompile(`'{2,5}`)
	regInternalLink  := regexp.MustCompile(`\[\[.*?\]\]`)
	regInternalLink1 := regexp.MustCompile(`(.*)\[\[(.*)#.*\]\](.*)`)
	regInternalLink2 := regexp.MustCompile(`(.*)\[\[(.*)\|.*\]\](.*)`)
	regExternalLink := regexp.MustCompile(`\[?http.*\]?`)
	regCommentOut := regexp.MustCompile(`<!--.*-->`)
	for _, v := range reg.FindAll([]byte(txt), -1) {
		s := string(v[1: len(v) - 2])
		strs := strings.Split(s, " = ")
		
		// 強調マークアップ
		tmp := regEmphasis.ReplaceAllString(strs[1], "")
		
		// 内部リンク
		if regInternalLink.FindString(tmp) == "" {
			m[strs[0]] = tmp
		} else if regInternalLink1.FindString(tmp) != "" {
			for _, s := range regInternalLink1.FindAllStringSubmatch(tmp, -1) {
				var res string
				for i := 1; i < len(s); i++ {
					res += s[i]
				}
				res = strings.Replace(res, "[[", "", -1)
				res = strings.Replace(res, "]]", "", -1)
				m[strs[0]] = res
			}
		} else if regInternalLink2.FindString(tmp) != "" {
			for _, s := range regInternalLink2.FindAllStringSubmatch(tmp, -1) {
				var res string
				for i := 1; i < len(s); i++ {
					res += s[i]
				}
				res = strings.Replace(res, "[[", "", -1)
				res = strings.Replace(res, "]]", "", -1)
				m[strs[0]] = res
			}
		} else {
			res := tmp
			res = strings.Replace(res, "[[", "", -1)
			res = strings.Replace(res, "]]", "", -1)
			m[strs[0]] = string(res)
		}
		// 外部リンク
		m[strs[0]] = regExternalLink.ReplaceAllString(m[strs[0]], "")		
		
		// コメントアウト
		m[strs[0]] = regCommentOut.ReplaceAllString(m[strs[0]], "")

	}
	
	for _, t := range(m) {
		fmt.Println(t)
	}
	
}
