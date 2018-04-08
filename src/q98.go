package main

import (
	"math"

	"github.com/cipepser/goClustering/vis"
	"github.com/cipepser/goClustering/ward"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"

	"encoding/gob"
	"log"
	"os"
)

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

	X := make([][]float64, len(vCountries))
	names := make([]string, len(vCountries))
	i := 0
	for c, vec := range vCountries {
		X[i] = vec
		names[i] = c
		i++
	}

	T := ward.Ward(X)

	d, err := vis.NewDendrogram(T)
	if err != nil {
		panic(err)
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Add(d)

	p.X.Label.Text = "Countories"
	p.NominalX(names...)

	p.X.Tick.Label.Rotation = math.Pi / 3
	p.X.Tick.Label.YAlign = draw.YCenter
	p.X.Tick.Label.XAlign = draw.XRight

	p.Y.Label.Text = "distance"

	file := "../data/q98.out.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}
}
