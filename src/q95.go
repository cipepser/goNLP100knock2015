package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	onlinestats "github.com/dgryski/go-onlinestats"
)

func main() {
	// f, err := os.Open("../data/q94_model_q85.out.txt")
	f, err := os.Open("../data/q94_model_q90.out.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReaderSize(f, 4096)

	// read first line to skip
	r.ReadLine()

	X := make([]float64, 0)
	Y := make([]float64, 0)

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		line := strings.Split(string(l), ",")
		x, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			panic(err)
		}
		X = append(X, x)

		y := -1.0
		if string(line[3]) != "NaN" {
			y, err = strconv.ParseFloat(line[3], 64)
			if err != nil {
				panic(err)
			}
		}
		Y = append(Y, y)
	}
	rs, _ := onlinestats.Spearman(X, Y)
	fmt.Println(rs)
}
