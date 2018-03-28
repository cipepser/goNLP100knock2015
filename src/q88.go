package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"os"
	"sort"

	"gonum.org/v1/gonum/mat"
)

type CosineSimilarity struct {
	word  string
	value float64
}

func main() {
	frp, err := os.Open("../data/q85.proj.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer frp.Close()

	dec := gob.NewDecoder(frp)

	var proj mat.Dense
	err = dec.Decode(&proj)
	if err != nil {
		log.Fatal(err)
	}

	frd, err := os.Open("../data/q85.dictt.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer frd.Close()

	dec = gob.NewDecoder(frd)

	dict := make(map[string]int)
	err = dec.Decode(&dict)
	if err != nil {
		log.Fatal(err)
	}

	idxEngland := dict["England"]
	// idxEngland := dict["United_States"] // 今のPCA結果では、U.Sが零ベクトルのため代理
	vEngland := proj.RowView(idxEngland)
	fmt.Println(vEngland)

	css := make([]CosineSimilarity, 0)
	for w, i := range dict {
		if i == idxEngland {
			continue
		}

		v := proj.RowView(i)
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
}

func cos(a, b mat.Vector) float64 {
	return mat.Dot(a, b) / (math.Sqrt(mat.Dot(a, a)) * math.Sqrt(mat.Dot(b, b)))
}
