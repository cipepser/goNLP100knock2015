package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	dim = 300
)

func main() {
	f, err := os.Open("../data/q92_model_q85.out.txt")
	// f, err := os.Open("../data/q92_model_q90.out.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReaderSize(f, 4096)

	correct := 0
	count := 0
	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if l[0] == ':' {
			continue
		}

		words := strings.Split(string(l), " ")

		if words[3] == words[4] {
			correct++
		}
		count++
	}

	fmt.Println(float64(correct) / float64(count))
}
