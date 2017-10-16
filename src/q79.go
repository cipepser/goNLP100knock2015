package main

import (
	"bufio"
	"errors"
	"image/color"
	"io"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strings"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"./q71"

	porterstemmer "github.com/reiver/go-porterstemmer"
)

func getIndexOfDict(s string, dict []string) int {
	for i, w := range dict {
		if w == s {
			return i
		}
	}
	return -1
}

func Sigmoid(v float64) float64 { return 1.0 / (1.0 + math.Exp(-v)) }

type Review struct {
	sentence []string
	label    string
}

func NewReview(l, s string) Review {
	r := Review{
		label: l,
	}
	r.sentence = strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-' || r == '/' || r == ';'
	})

	return r
}

func PreProcessing(rs []Review) []Review {
	for i := range rs {
		rs[i].sentence = RemoveStopWords(rs[i].sentence)
		rs[i].sentence = RemoveMetaCharacters(rs[i].sentence)
		rs[i].sentence = Stemming(rs[i].sentence)
	}

	return rs
}

func RemoveMetaCharacters(input []string) (output []string) {
	metas := []string{
		"`", "[", "]", ";", ":", "/", "*",
		"(", ")", ".", "\"", ",", "&", "?",
		"!", "%", "'",
	}
	//     s = strings.Replace(s, "'", "", -1)

	for _, s := range input {
		for _, m := range metas {
			s = strings.Replace(s, m, "", -1)
		}
		if s != "" {
			output = append(output, s)
		}
	}

	return output
}

func RemoveStopWords(input []string) (output []string) {
	for _, s := range input {
		if !q71.IsStopWords(s) {
			output = append(output, s)
		}
	}

	return output
}

func Stemming(str []string) []string {
	for i := range str {
		str[i] = porterstemmer.StemString(str[i])
	}

	return str
}

func makeDictionary(rs []Review) map[string]int {
	dict := make(map[string]int, 0)

	for _, r := range rs {
		for _, w := range r.sentence {
			dict[w]++
		}
	}

	// remove noise
	// for w := range dict {
	// 	if dict[w] < 6 {
	// 		delete(dict, w)
	// 	}
	// }

	return dict
}

func makeFeatureVectors(rs []Review, dict map[string]int) ([]*mat.VecDense, []string) {
	feature := make([]string, len(dict)+1)
	i := 0
	for w := range dict {
		feature[i] = w
		i++
	}

	X := make([]*mat.VecDense, len(rs))
	for i, r := range rs {
		x := make([]float64, len(feature))

		for _, w := range r.sentence {
			idx := getIndexOfDict(w, feature)
			if idx != -1 {
				x[idx] = 1
			}
		}
		x[len(x)-1] = 1
		X[i] = mat.NewVecDense(len(x), x)
	}

	return X, feature
}

// LogisticRegression returns w which is the weight vector by logistic regressin.
func LogisticRegression(X []*mat.VecDense, labels []string, eta float64) (*mat.VecDense, error) {
	if len(X) != len(labels) {
		return nil, errors.New("X and label must have same length.")
	}

	// correct label
	t := make([]float64, len(X))
	for i, l := range labels {
		if l == "+1" {
			t[i] = 1.0
		}
	}
	// initialize the parameter
	ws := make([]float64, X[0].Len())
	for i := range ws {
		ws[i] = rand.Float64()
	}
	w := mat.NewVecDense(len(ws), ws)

	// training
	for i := range X {
		x := mat.NewVecDense(X[i].Len(), nil)
		x.CopyVec(X[i])

		p := Sigmoid(mat.Dot(w, x))

		x.ScaleVec(eta*(p-t[i]), x)
		w.SubVec(w, x)

		eta *= 0.99999
	}

	return w, nil
}

func Predict(w, x *mat.VecDense, th float64) (string, float64) {
	p := Sigmoid(mat.Dot(w, x))

	if p > th {
		return "+1", p
	}
	return "-1", p
}

type Feature struct {
	word   string
	weight float64
}

func myScatter(x, y []float64) {
	d := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		d[i].X = x[i]
		d[i].Y = y[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "precision - recall"
	p.X.Label.Text = "precision"
	p.Y.Label.Text = "recall"
	s, err := plotter.NewScatter(d)
	if err != nil {
		panic(err)
	}
	s.Radius = vg.Length(2)
	s.Color = color.RGBA{R: 255, B: 0, A: 255}
	p.Add(s)

	file := "precison_vs_recall.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}

func myScatters(x, y1, y2 []float64) {
	d1 := make(plotter.XYs, len(x))
	d2 := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		d1[i].X = x[i]
		d1[i].Y = y1[i]

		d2[i].X = x[i]
		d2[i].Y = y2[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "threshold - precision, recall"
	p.X.Label.Text = "threshold"
	p.Y.Label.Text = "precision, recall"
	s1, err := plotter.NewScatter(d1)
	if err != nil {
		panic(err)
	}
	s1.Radius = vg.Length(2)
	s1.Color = color.RGBA{R: 255, B: 0, A: 255}
	p.Add(s1)
	p.Legend.Add("precision", s1)

	s2, err := plotter.NewScatter(d2)
	if err != nil {
		panic(err)
	}
	s2.Radius = vg.Length(2)
	s2.Color = color.RGBA{R: 0, B: 255, A: 255}
	p.Add(s2)
	p.Legend.Add("recall", s2)
	p.Legend.Top = true

	file := "threshold_vs_precison_and_recall.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}

func main() {
	rs := []Review{}

	f, err := os.Open("../data/sentiment.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// split to label and sentence.
		str := strings.SplitN(string(l), " ", 2)
		r := NewReview(str[0], str[1])
		rs = append(rs, r)
	}

	rs = PreProcessing(rs)
	dict := makeDictionary(rs)
	X, _ := makeFeatureVectors(rs, dict)

	labels := make([]string, len(rs))
	for i, r := range rs {
		labels[i] = r.label
	}

	eta := 0.6
	w, err := LogisticRegression(X, labels, eta)
	if err != nil {
		panic(err)
	}

	ths := []float64{0, 0.5, 0.1, 0.15, 0.2, 0.25, 0.3, 0.35, 0.4, 0.45}
	pres := []float64{}
	recs := []float64{}
	for _, th := range ths {
		correct := 0
		actPos := 0
		prePos := 0
		andPos := 0
		for i, x := range X {
			ans, _ := Predict(w, x, th)

			if ans == labels[i] {
				correct++
			}

			if labels[i] == "+1" {
				actPos++
			}
			if ans == "+1" {
				prePos++
			}

			if labels[i] == "+1" && ans == "+1" {
				andPos++
			}
		}

		preRate := float64(andPos) / float64(prePos)
		pres = append(pres, preRate)

		recRate := float64(andPos) / float64(actPos)
		recs = append(recs, recRate)
	}

	myScatters(ths, pres, recs)
	myScatter(pres, recs)
}
