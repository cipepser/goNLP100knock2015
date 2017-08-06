package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/reiver/go-porterstemmer"
)

func main() {
	f, err := os.Open("../data/q51_out.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)

	for {
		w, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		if len(w) == 0 {
			continue
		}

		stem := porterstemmer.StemString(string(w))
		fmt.Println(string(w), "\t", stem)
	}
}
