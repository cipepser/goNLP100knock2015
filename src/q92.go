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
	f, err := os.Open("../data/q91.out.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReaderSize(f, 4096)

	// model := loadModel("../data/q85_model.txt", dim)
	model := loadModel("../data/trained_model.txt", dim)

	// fw, err := os.Create("../data/q92_model_q85.out.txt")
	fw, err := os.Create("../data/q92_model_q90.out.txt")
	defer fw.Close()
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriterSize(fw, 4096)
	defer w.Flush()

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		w.Write(l)

		if l[0] == ':' {
			w.Write([]byte("\n"))
			continue
		}

		words := strings.Split(string(l), " ")

		vec := make([]float64, dim)
		if len(model[words[0]]) != dim || len(model[words[1]]) != dim || len(model[words[2]]) != dim {
			w.Write([]byte(" Undifined NaN\n"))
		} else {
			floats.Add(vec, model[words[1]])
			floats.Sub(vec, model[words[0]])
			floats.Add(vec, model[words[2]])

			word, cosine := getBestSimilarity(model, vec)
			w.Write([]byte(" " + word + " " + strconv.FormatFloat(cosine, 'f', 6, 64) + "\n"))
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

func getBestSimilarity(model map[string][]float64, vec []float64) (word string, cosine float64) {
	cosine = -1.1
	word = "Undifined"

	for w, v := range model {
		c := cos(vec, v)
		if c > cosine {
			cosine = c
			word = w
		}
	}

	return word, cosine
}
