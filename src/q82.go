package main

import (
	"bufio"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("../data/q81.out.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	corpus := []string{}
	r := bufio.NewReaderSize(f, 4096)
	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		tokens := strings.Split(string(l), " ")
		for _, t := range tokens {
			corpus = append(corpus, t)
		}
	}

	fw, err := os.Create("../data/q82.out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fw.Close()

	rand.Seed(time.Now().UnixNano())
	for i, c := range corpus {
		d := rand.Intn(5) + 1
		for j := 1; j <= d; j++ {
			if i-j+1 > 0 {
				fw.WriteString(c + "\t" + corpus[i-j])
				fw.WriteString("\n")
			}
		}
		for j := 1; j <= d; j++ {
			if i+j < len(corpus) {
				fw.WriteString(c + "\t" + corpus[i+j])
				fw.WriteString("\n")
			}
		}
	}
}
