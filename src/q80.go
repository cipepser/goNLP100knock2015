package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func removeSymbol(s string) string {
	if len(s) == 0 {
		return s
	}

	noise := []string{".", ",", "!", "?", ";", ":", "(", ")", "[", "]", "'", "\""}

	for _, v := range noise {
		if string(s[0]) == v {
			s = strings.TrimPrefix(s, v)
			s = removeSymbol(s)
			if len(s) == 0 {
				return s
			}
		}
		if string(s[len(s)-1]) == v {
			s = strings.TrimSuffix(s, v)
			s = removeSymbol(s)
			if len(s) == 0 {
				return s
			}
		}
	}
	return s
}

func main() {
	f, err := os.Open("../data/enwiki-20150112-400-r100-10576.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	corpus := [][]string{}
	r := bufio.NewReaderSize(f, 4096)
	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		c := []string{}
		tokens := strings.Split(string(l), " ")
		for _, t := range tokens {
			t = removeSymbol(t)
			if len(t) > 1 {
				c = append(c, t)
			}
		}
		if len(c) > 0 {
			corpus = append(corpus, c)
		}
	}

	// write the result to txt file
	fw, err := os.Create("../data/q80.out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fw.Close()

	for _, ts := range corpus {
		for i, t := range ts {
			fw.Write([]byte(t))
			if i != len(ts)-1 {
				fw.Write([]byte(string(' ')))
			}
		}
		fw.Write([]byte("\n"))
	}

}
