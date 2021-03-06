package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"sort"
	"strings"

	"gonum.org/v1/gonum/mat"

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

func Predict(w, x *mat.VecDense) (string, float64) {
	p := Sigmoid(mat.Dot(w, x))

	if p > 0.5 {
		return "+1", p
	}
	return "-1", p
}

type Feature struct {
	word   string
	weight float64
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

	// q72
	rs = PreProcessing(rs)
	dict := makeDictionary(rs)
	X, feature := makeFeatureVectors(rs, dict)

	// q73
	labels := make([]string, len(rs))
	for i, r := range rs {
		labels[i] = r.label
	}

	eta := 0.6
	w, err := LogisticRegression(X, labels, eta)
	if err != nil {
		panic(err)
	}

	// q74, 76
	fmt.Println("**** labeling and the probablity ****")
	fmt.Println("No.\tcorrect\tpredicted\tprobability")
	fmt.Println("--------------------------------------------")
	for i, x := range X[:10] {
		ans, p := Predict(w, x)
		fmt.Println("[", i, "]\t", labels[i], "\t", ans, "\t", p)
	}
	fmt.Println("")

	// q75
	features := make([]Feature, w.Len())
	for i := 0; i < w.Len(); i++ {
		features[i].weight = w.At(i, 0)
		features[i].word = feature[i]
	}

	sort.Slice(features, func(i, j int) bool {
		return features[i].weight < features[j].weight
	})

	fmt.Println("**** 10 highest weighted words ****")
	fmt.Println("rank\tword\tweight")
	fmt.Println("--------------------------------------------")
	for i := 0; i < 10; i++ {
		fmt.Println("[", i, "]\t", features[len(features)-i-1].word,
			"\t", features[len(features)-i-1].weight)
	}
	fmt.Println("")
	fmt.Println("**** 10 lowest weighted words ****")
	fmt.Println("rank\tword\tweight")
	fmt.Println("--------------------------------------------")
	for i := 0; i < 10; i++ {
		fmt.Println("[", i, "]\t", features[i].word, "\t", features[i].weight)
	}
	fmt.Println("")

	// q77
	correct := 0
	actPos := 0
	prePos := 0
	andPos := 0
	for i, x := range X {
		ans, _ := Predict(w, x)

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

	fmt.Println("**** rates ****")
	fmt.Println("accuracy rate:\t", float64(correct)/float64(len(X)))

	preRate := float64(andPos) / float64(prePos)
	fmt.Println("precision rate:\t", preRate)

	recRate := float64(andPos) / float64(actPos)
	fmt.Println("recall rate:\t", recRate)

	fmt.Println("F1 score:\t", 2*(preRate*recRate)/(preRate+recRate))
}
