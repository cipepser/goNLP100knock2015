package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"os"

	"gonum.org/v1/gonum/mat"
)

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

	v1 := proj.RowView(dict["United_States"])
	// v2 := proj.RowView(dict["U.S"])
	v2 := proj.RowView(94) // 今のPCA結果では、U.Sが零ベクトルのため代理

	fmt.Println(cos(v1, v2))
}

func cos(a, b mat.Vector) float64 {
	return mat.Dot(a, b) / (math.Sqrt(mat.Dot(a, a)) * math.Sqrt(mat.Dot(b, b)))
}
