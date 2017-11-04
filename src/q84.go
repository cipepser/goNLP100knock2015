package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
)

type key struct {
	t, c string
}

type kv struct {
	k key
	v float64
}

func calcPPMI(N, tc, t, c int) float64 {
	if tc < 10 {
		return 0.0
	}

	v := 0.0
	v = math.Log(float64(N*tc) / float64(t*c))

	if v > 0 {
		return v
	}
	return 0.0
}

func main() {
	f, err := os.Open("../data/q82.out.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReaderSize(f, 4096)
	m := make(map[key]int)
	Nt := make(map[string]int)
	Nc := make(map[string]int)

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		str := strings.Split(string(l), "\t")
		k := key{
			t: str[0],
			c: str[1],
		}
		m[k]++
		Nt[k.t]++
		Nc[k.c]++
	}

	X := make(map[key]float64)
	for k := range m {
		X[k] = calcPPMI(len(m), m[k], Nt[k.t], Nc[k.t])
	}

	fmt.Println("the number of the word: ", len(Nt))
	fmt.Println("the number of contex word: ", len(Nc))
	fmt.Println("length of X: ", len(X))
	fmt.Println("n(X)/(n(t) * n(c)): ", float64(len(X))/float64(len(Nt)*len(Nc)))
	fmt.Println("------------------")

	Xs := make([]kv, len(X))
	i := 0
	for k, v := range X {
		Xs[i] = kv{k, v}
		i++
	}

	sort.Slice(Xs, func(i, j int) bool {
		return Xs[i].v > Xs[j].v
	})

	for _, v := range Xs[:10] {
		fmt.Println(v)
	}

}
