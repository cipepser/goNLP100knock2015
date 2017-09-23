package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	sentiment := []string{}

	// label for positive sentence.
	fpos, err := os.Open("../data/rt-polaritydata/rt-polarity.pos")
	defer fpos.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReaderSize(fpos, 4096)
	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		sentiment = append(sentiment, "+1 "+string(l))
	}

	// label for negative sentence.
	fneg, err := os.Open("../data/rt-polaritydata/rt-polarity.neg")
	defer fneg.Close()
	if err != nil {
		panic(err)
	}

	r = bufio.NewReaderSize(fneg, 4096)
	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		sentiment = append(sentiment, "-1 "+string(l))
	}

	// shuffle sentences
	sent := make([]string, len(sentiment))
	rand.Seed(time.Now().UnixNano())
	for i, n := range rand.Perm(len(sentiment)) {
		sent[i] = sentiment[n]
	}

	// write the result to txt file
	fw, err := os.Create("../data/sentiment.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fw.Close()

	for _, s := range sent {
		fw.Write([]byte(s))
		fw.Write([]byte(string('\n')))
	}

	// count the number of positive sentences and negative.
	var p, n int
	f, err := os.Open("../data/sentiment.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r = bufio.NewReaderSize(f, 4096)
	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		s := strings.Split(string(l), " ")[0]
		if s == "+1" {
			p++
		} else {
			n++
		}
	}

	fmt.Println("positive: ", p)
	fmt.Println("negative: ", n)
}
