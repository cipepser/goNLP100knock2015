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

	// 94, 130, 249 // 今のPCA結果では、零ベクトルのため代理
	vSpain := mat.NewVecDense(len(proj.RawRowView(dict["Spain"])), proj.RawRowView(dict["Spain"]))
	vMadrid := mat.NewVecDense(len(proj.RawRowView(dict["Madrid"])), proj.RawRowView(dict["Madrid"]))
	vAthens := mat.NewVecDense(len(proj.RawRowView(dict["Athens"])), proj.RawRowView(dict["Athens"]))
	// vSpain := mat.NewVecDense(len(proj.RawRowView(94)), proj.RawRowView(94))
	// vMadrid := mat.NewVecDense(len(proj.RawRowView(130)), proj.RawRowView(130))
	// vAthens := mat.NewVecDense(len(proj.RawRowView(249)), proj.RawRowView(249))

	vTarget := new(mat.VecDense)
	vTarget.SubVec(vSpain, vMadrid)
	vTarget.AddVec(vTarget, vAthens)

	css := make([]CosineSimilarity, 0)
	for w, i := range dict {
		v := proj.RowView(i)
		cosine := cos(vTarget, v)

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
