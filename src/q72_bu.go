package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	porterstemmer "github.com/reiver/go-porterstemmer"

	"gonum.org/v1/gonum/mat"

	"./q71"
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

func main() {
	f, err := os.Open("../data/sentiment.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	labels := []string{}
	sentences := [][]string{}
	dict := make(map[string]int, 0)

	// make a dictionary
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
		// n, err := strconv.Atoi(str[0])
		// if err != nil {
		// 	panic(err)
		// }

		labels = append(labels, str[0])

		// tmp := strings.Split(str[1], " ")
		tmp := strings.FieldsFunc(str[1], func(r rune) bool {
			return r == ' ' || r == '-' || r == '/' || r == ';'
		})

		sent := []string{}
		metas := []string{"`", "[", "]", ";", ":", "/", "*", "(", ")", ".", "\"", ",", "&", "?"}
		for _, s := range tmp {
			// trim meta charters
			for _, m := range metas {
				s = strings.Replace(s, m, "", -1)
			}

			// remove stop words
			if !q71.IsStopWords(s) {
				s = porterstemmer.StemString(s) // stemming
				s = strings.Replace(s, "'", "", -1)
				if s != "" {
					sent = append(sent, s)
					dict[s]++
				}
			}
		}
		if sent != nil {
			sentences = append(sentences, sent)
			// fmt.Println(sent)
		}
	}

	// remove noise
	fmt.Println(len(dict))
	for w := range dict {
		if dict[w] < 6 {
			delete(dict, w)
		}
		// if dict[w] > 20000 {
		// 	delete(dict, w)
		// }
	}
	fmt.Println(len(dict))

	// fmt.Println(dict)

	feature := make([]string, len(dict)+1)
	i := 0
	for w := range dict {
		feature[i] = w
		i++
	}

	// make feature vectors
	X := make([][]float64, len(sentences))
	for i, sent := range sentences {
		x := make([]float64, len(feature))

		for _, s := range sent {
			idx := getIndexOfDict(s, feature)
			// if idx == -1 {
			// 	panic(s + " is not found in the dictionary.")
			// }
			if idx != -1 {
				x[idx] = 1
				// x[idx]++
			}
		}
		x[len(x)-1] = 1
		X[i] = x
		// fmt.Println(x)
	}

	// correct label
	t := make([]float64, len(X))
	for i, l := range labels {
		if l == "+1" {
			t[i] = 1.0
		}
	}
	// fmt.Println(t)

	// initialize the parameter
	ws := make([]float64, len(feature))
	// for i := range ws {
	// 	ws[i] = rand.Float64()
	// }
	w := mat.NewVecDense(len(ws), ws)

	// training
	eta := 0.6
	for i := range X {
		x := mat.NewVecDense(len(X[i]), X[i])

		p := Sigmoid(mat.Dot(w, x))

		// fmt.Println("----------")
		// fmt.Println("eta:", eta)
		// fmt.Println("p: ", p)
		// fmt.Println("t: ", t[i])
		// fmt.Println("eta*(p-t): ", eta*(p-t[i]))

		x.ScaleVec(eta*(p-t[i]), x)
		w.SubVec(w, x)

		eta *= 0.99999
	}

	// for i := 0; i < w.Len(); i++ {
	// 	// if w.At(i, 0) > 0 {
	// 	fmt.Println(w.At(i, 0))
	// 	// }
	// }

	// classification
	correct := 0
	for i := range X {
		x := mat.NewVecDense(len(X[i]), X[i])

		// fmt.Println(mat.Dot(w, x))
		p := Sigmoid(mat.Dot(w, x))

		label := 0.0
		if p > 0.5 {
			label = 1.0
		}
		// fmt.Println("t: ", t[i])
		// fmt.Println("label: ", label)
		// fmt.Println("p: ", p)
		// fmt.Println("-----------")
		if t[i] == label {
			correct++
		}
	}
	fmt.Println(float64(correct) / float64(len(X)))
	// fmt.Println("---------------")
	// for i := 0; i < len(feature); i++ {
	// 	fmt.Println(feature[i], w.At(i, 0))
	// }

	// x := mat.NewVecDense(len(X[i]), X[i])
	// p := Sigmoid(mat.Dot(w, x))

}
