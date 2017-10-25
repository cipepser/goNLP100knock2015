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
		for string(s[0]) == v {
			s = strings.TrimPrefix(s, v)
			if len(s) == 0 {
				return s
			}
		}
		// fmt.Println(s)
		for string(s[len(s)-1]) == v {
			s = strings.TrimSuffix(s, v)
			if len(s) == 0 {
				return s
			}
		}
	}
	return s
}

func main() {
	f, err := os.Open("../data/enwiki-20150112-400-r100-10576.tmp.txt")
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

	// for _, c := range corpus {
	// 	fmt.Println(c)
	// }

	// write the result to txt file
	fw, err := os.Create("../data/q80.out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fw.Close()

	for _, ts := range corpus {
		for _, t := range ts {
			fw.Write([]byte(t))
			fw.Write([]byte(string(' ')))
		}
		fw.Write([]byte("\n"))
	}

}
