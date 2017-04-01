package main

import (
	"os"
	"io"
	"bufio"
	"strings"
	"sort"
	"image/color"
	"math"
	"errors"
	"log"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}
func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}
func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

// sortedKeys returns the slice []string which is sorted by the count in map.
// Eventually, we should hold the sorted slice []string, it's enough to sort.
// func (sm *sortedMap) sortedKeys(m map[string]int) []string {
func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	
	sort.Sort(sm)
	
	return sm.s
}

// plotScatter plots a scatter of X and Y.
// X and Y SHOULD have same length.
func plotScatter(X, Y []float64) error {
	if len(X) != len(Y) {
		return errors.New("X and Y SHOULD have same length.")
	}

	scatterData := make(plotter.XYs, len(X))
	for i, _ := range X {
		scatterData[i].X = X[i]
		scatterData[i].Y = Y[i]
	}

	p, err := plot.New()
	if err != nil {
		 panic(err)
	}
	p.Title.Text = "Zipf's law"
	p.X.Label.Text = "common logarithm of the rank of frequency"
	p.Y.Label.Text = "common logarithm of the frequency"

	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(scatterData)
	if err != nil {
		 panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	s.GlyphStyle.Radius = vg.Points(1)

	p.Add(s)
	p.Legend.Add("scatter", s)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "q39.png"); err != nil {
		 panic(err)
	}
	
	return nil
}



func main()  {
	f, err := os.Open("../data/neko.txt.mecab")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	
	r := bufio.NewReader(f)	
	
	sents := make([][]map[string]string, 0)
	sent := make([]map[string]string, 0)

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		
		// store morpheme which is not "EOS" into maps
		if string(b) != "EOS" {
			// split by tab and comma
			tmp := strings.Split(string(b), "\t")
			m := append(tmp[:1], strings.Split(tmp[1], ",")...)
			
			morpheme := make(map[string]string)
			
		 	morpheme["surface"] = m[0]
			morpheme["base"]    = m[7]
			morpheme["pos"]     = m[1]
			morpheme["pos1"]    = m[2]
			
			sent = append(sent, morpheme)
		} else { // if we find "EOS", store sentence to sentences and initialize the sent
			if len(sent) > 0 { // for appearing "EOS" continuously
				sents = append(sents, sent)
				sent = make([]map[string]string, 0)
			}
		}
	}
	
	// count the number of the morpheme has same base
	freq := make(map[string]int)
	for _, sent := range sents {
		for _, m := range sent {
			freq[m["base"]]++
		}
	}
	
	// draw the bi-logarithm graph
	res := sortedKeys(freq)	
	var X, Y []float64


	for i, v := range res {
		X = append(X, math.Log10(float64(i + 1)))
		Y = append(Y, math.Log10(float64(freq[v])))
	}
	
	err = plotScatter(X, Y)
	if err != nil {
		log.Fatal(err)
	}
	
}
