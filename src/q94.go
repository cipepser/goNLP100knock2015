package main

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/floats"
)

var (
	dim = 300
)

func main() {
	f, err := os.Open("../data/wordsim353/combined.csv")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReaderSize(f, 4096)

	model := loadModel("../data/q85_model.txt", dim)
	// model := loadModel("../data/trained_model.txt", dim)

	fw, err := os.Create("../data/q94_model_q85.out.txt")
	// fw, err := os.Create("../data/q94_model_q90.out.txt")
	defer fw.Close()
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriterSize(fw, 4096)
	defer w.Flush()

	l, _, _ := r.ReadLine()
	w.Write(l)
	w.Write([]byte("\n"))

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		w.Write(l)

		words := strings.Split(string(l), ",")

		if len(model[words[0]]) != dim || len(model[words[1]]) != dim {
			w.Write([]byte(",NaN\n"))
		} else {
			w.Write([]byte("," + strconv.FormatFloat(cos(model[words[0]], model[words[1]]), 'f', 6, 64) + "\n"))
		}
	}
}

func loadModel(file string, dim int) map[string][]float64 {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReaderSize(f, 4096)

	vec := make(map[string][]float64, 0)

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		str := strings.Split(string(l), " ")
		if len(str) != dim+2 {
			continue
		}

		data := make([]float64, dim)
		for i := 1; i < dim+1; i++ {
			data[i-1], err = strconv.ParseFloat(str[i], 64)
			if err != nil {
				panic(err)
			}
		}
		vec[str[0]] = data
	}
	return vec
}

func cos(a, b []float64) float64 {
	return floats.Dot(a, b) / (math.Sqrt(floats.Dot(a, a)) * math.Sqrt(floats.Dot(b, b)))
}
