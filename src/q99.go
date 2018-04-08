package main

import (
	"encoding/gob"
	"image/color"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"time"

	"./q99"

	"github.com/sacado/tsne4go"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
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

	X := make(tsne4go.VectorDistancer, len(Countries))
	names := make([]string, len(Countries))
	colors := make([]color.Color, len(Countries))
	for i, c := range Countries {
		X[i] = c.vector
		names[i] = c.name
		switch c.label {
		case 1:
			colors[i] = color.RGBA{R: 255, G: 0, B: 0, A: 255}
		case 2:
			colors[i] = color.RGBA{R: 0, G: 0, B: 255, A: 255}
		case 3:
			colors[i] = color.RGBA{R: 60, G: 179, B: 113, A: 255}
		case 4:
			colors[i] = color.RGBA{R: 238, G: 130, B: 238, A: 255}
		case 5:
			colors[i] = color.RGBA{R: 255, G: 165, B: 0, A: 255}
		}

	}

	t := tsne4go.New(X, nil)

	e := 1.0
	for i := 0; i < 500; i++ {
		enew := t.Step()
		if math.Abs(e-enew) < 1e-10 {
			break
		}
		e = enew
	}

	Y := t.Solution
	MyScatter(Y, names, colors)
}

func MyScatter(X []tsne4go.Point, names []string, colors []color.Color) {
	x := make([]float64, len(X))
	y := make([]float64, len(X))
	cs := make([]color.Color, len(X))
	for i := 0; i < len(X); i++ {
		x[i] = X[i][0]
		y[i] = X[i][1]
		cs[i] = colors[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	s, err := q99.NewScatter(x, y, names, cs)
	if err != nil {
		panic(err)
	}

	p.Add(s)

	file := "q99.out.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
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
