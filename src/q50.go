package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	f, err := os.Open("../data/nlp.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)

	rTitle := regexp.MustCompile(`.+`)
	rSent := regexp.MustCompile(`.*?(\.|;|:|\?|!)\s[A-Z]{1}?`)

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		// add dummy character to capture the last sentence.
		l = []byte(string(l) + " A")

		// capture the sentecse.
		flg := false
		var capital byte
		for i, v := range rSent.FindAll(l, -1) {
			if i != len(v)-1 {
				v = append([]byte{capital}, v...)
				capital = v[len(v)-1]
			}
			v = v[:len(v)-2]

			fmt.Println(string(v))
			flg = true
		}
		if flg {
			continue
		}

		// capture the title.
		for _, v := range rTitle.FindAll(l, -1) {
			// remove dummy character.
			v = v[:len(v)-2]
			fmt.Println(string(v))
		}
	}
}
