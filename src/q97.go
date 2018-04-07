package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"time"

	"gonum.org/v1/gonum/floats"
)

var (
	dim = 300
	k   = 5
)

type Country struct {
	name   string
	vector []float64
	label  int
}

func main() {
	f, err := os.Open("../data/q96.out.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := gob.NewDecoder(f)
	var vCountries map[string][]float64
	err = dec.Decode(&vCountries)
	if err != nil {
		log.Fatal(err)
	}

	Countries := kMeans(vCountries)

	for _, c := range Countries {
		fmt.Println(c.label, ":", c.name)
	}
}

func kMeans(vCountries map[string][]float64) []Country {
	Countries := initialize(vCountries)

	for {
		V := calcCenters(Countries)
		cs := relabeling(Countries, V)
		if reflect.DeepEqual(Countries, cs) {
			break
		}
		Countries = cs
	}

	return Countries
}

func initialize(vCountries map[string][]float64) []Country {
	rand.Seed(time.Now().UnixNano())

	Countries := []Country{}
	for c, vec := range vCountries {
		Countries = append(Countries, Country{
			name:   c,
			vector: vec,
			label:  rand.Intn(k) + 1,
		})
	}

	return Countries
}

func calcCenters(Countries []Country) [][]float64 {
	V := make([][]float64, k)
	for i := range V {
		V[i] = make([]float64, dim)
	}
	N := make([]int, k)

	for _, c := range Countries {
		floats.Add(V[c.label-1], c.vector)
		N[c.label-1]++
	}

	for i := range V {
		for j := range V[i] {
			V[i][j] = V[i][j] / float64(N[i])
		}
	}

	return V
}

func relabeling(Countries []Country, V [][]float64) []Country {
	for i, c := range Countries {
		min := math.Inf(1)
		for j, v := range V {
			d := floats.Distance(v, c.vector, 2)
			if min > d {
				min = d
				Countries[i].label = j + 1
			}
		}
	}
	return Countries
}
