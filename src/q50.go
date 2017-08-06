package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	f, err := os.Open("../data/nlp_tmp.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)

	// (. or ; or : or ? or !) → 空白文字 → 英大文字というパターンを文の区切りと見なし，
	// 入力された文書を1行1文の形式で出力せよ．
	rTitle := regexp.MustCompile(`.+`)
	// rSent := regexp.MustCompile(`.*(\.|;|:|\?|!)\s(A-Z)*`)
	rSent := regexp.MustCompile(`.*?(\.|;|:|\?|!)\s(A-Z)*`)
	// rSent := regexp.MustCompile(`(.*[\.;:?!])\s(A-Z)*`)

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		// to capture the last sentence.
		l = []byte(string(l) + "\n")

		// capture the sentecse
		flg := false
		for _, v := range rSent.FindAll(l, -1) {
			fmt.Println("---------------")

			// remove the LF
			if string(v[len(v)-1]) == "\n" {
				v = v[:len(v)-1]
			}
			fmt.Println(string(v))
			flg = true
		}
		if flg {
			continue
		}

		// capture the title.
		for _, v := range rTitle.FindAll(l, -1) {
			fmt.Println("---------------")

			fmt.Println(string(v))
		}
	}
}
