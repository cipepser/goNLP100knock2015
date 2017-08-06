package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../data/q50_out.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)

	for {
		s, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		for _, w := range strings.Split(string(s), " ") {
			fmt.Println(w)
		}
		fmt.Println("")
	}
}
