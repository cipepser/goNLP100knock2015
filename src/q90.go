package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/floats"
)

var (
	dim = 300
)

type CosineSimilarity struct {
	word  string
	value float64
}

func main() {
	f, err := os.Open("../data/trained_model.txt")
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
				log.Fatal(err)
			}
		}
		vec[str[0]] = data
	}

	fmt.Println("------------------------------")
	fmt.Println("86. 単語ベクトルの表示")
	fmt.Println("United_States")
	fmt.Println(vec["United_States"])

	fmt.Println("------------------------------")
	fmt.Println("87. 単語の類似度")
	fmt.Println("United_Stetes v.s. U.S.")
	fmt.Println(cos(vec["United_States"], vec["U.S"]))

	fmt.Println("------------------------------")
	fmt.Println("88. 類似度の高い単語10件")
	vEngland := vec["England"]
	css := make([]CosineSimilarity, 0)
	for w, v := range vec {
		cosine := cos(vEngland, v)
		if !math.IsNaN(cosine) {
			css = append(css, CosineSimilarity{
				word:  w,
				value: cosine,
			})
		}
	}

	sort.Slice(css, func(i, j int) bool {
		return css[i].value > css[j].value
	})

	for i, cs := range css {
		if i > 9 {
			break
		}
		fmt.Println(i+1, " ", cs.word, ":", cs.value)
	}

	fmt.Println("------------------------------")
	fmt.Println("89. 加法構成性によるアナロジー")
	vector := make([]float64, dim)
	floats.Add(vector, vec["Spain"])
	floats.Sub(vector, vec["Madrid"])
	floats.Add(vector, vec["Athens"])
	css = make([]CosineSimilarity, 0)
	for w, v := range vec {
		cosine := cos(vector, v)
		if !math.IsNaN(cosine) {
			css = append(css, CosineSimilarity{
				word:  w,
				value: cosine,
			})
		}
	}

	sort.Slice(css, func(i, j int) bool {
		return css[i].value > css[j].value
	})

	for i, cs := range css {
		if i > 9 {
			break
		}
		fmt.Println(i+1, " ", cs.word, ":", cs.value)
	}

}

func cos(a, b []float64) float64 {
	return floats.Dot(a, b) / (math.Sqrt(floats.Dot(a, a)) * math.Sqrt(floats.Dot(b, b)))
}
