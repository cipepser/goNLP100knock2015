package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/james-bowman/sparse"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

const (
	myTimeFormat = "2006/01/02 15:04:05 JST" // to display progress status
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
	fmt.Println("read file:", time.Now().Format(myTimeFormat))
	f, err := os.Open("../data/q82_tmp.out.txt")
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
	fmt.Println("\tfinished:", time.Now().Format(myTimeFormat))

	fmt.Println("calculate PPMI:", time.Now().Format(myTimeFormat))
	X := make(map[key]float64)
	for k := range m {
		ppmi := calcPPMI(len(m), m[k], Nt[k.t], Nc[k.c])
		if ppmi > 0 {
			X[k] = ppmi
		}
	}
	fmt.Println("\tfinished:", time.Now().Format(myTimeFormat))
	fmt.Println("\tthe number of the word:", len(Nt))
	fmt.Println("\tthe number of contex word: ", len(Nc))
	fmt.Println("\tfinished:", time.Now().Format(myTimeFormat))

	fmt.Println("transform hashmap to matrix:", time.Now().Format(myTimeFormat))
	fmt.Println("\tre-order idxt:", time.Now().Format(myTimeFormat))
	idxt := make(map[int]string)
	dictt := make(map[string]int)
	i := 0
	for k := range Nt {
		idxt[i] = k
		dictt[k] = i
		i++
	}

	fmt.Println("\tre-order idxc:", time.Now().Format(myTimeFormat))
	idxc := make(map[int]string)
	dictc := make(map[string]int)
	j := 0
	for k := range Nc {
		idxc[j] = k
		dictc[k] = j
		j++
	}

	fmt.Println("\tstore data as COO:", time.Now().Format(myTimeFormat))
	data := make([]float64, len(X))
	ia := make([]int, len(X))
	ja := make([]int, len(X))

	i = 0
	for k, v := range X {
		data[i] = v
		ia[i] = dictt[k.t]
		ja[i] = dictc[k.c]
		i++
	}
	// y := sparse.NewCSR(len(Nt), len(Nc), ia, ja, data)
	y := sparse.NewCOO(len(Nt), len(Nc), ia, ja, data)
	fmt.Println("\tfinished:", time.Now().Format(myTimeFormat))
	_ = y

	fmt.Println("PCA start:", time.Now().Format(myTimeFormat))

	var pc stat.PC
	// // TODO: NewDenseされてmakesliceのout of rangeになる
	ok := pc.PrincipalComponents(y, nil)
	if !ok {
		log.Fatal("PCA fails")
	}

	fmt.Println(time.Now().Format(myTimeFormat))
	fmt.Println("\tfinished:", time.Now().Format(myTimeFormat))

	fmt.Println("project to 300 dimensions:", time.Now().Format(myTimeFormat))
	k := 300
	var proj mat.Dense
	proj.Mul(y, pc.VectorsTo(nil).Slice(0, len(Nc), 0, k))
	fmt.Println("\tfinished:", time.Now().Format(myTimeFormat))

	fmt.Println("save proj:", time.Now().Format(myTimeFormat))
	fwp, err := os.Create("../data/q85.proj.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fwp.Close()

	enc := gob.NewEncoder(fwp)
	err = enc.Encode(proj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\tfinished:", time.Now().Format(myTimeFormat))

	fmt.Println("save dictt:", time.Now().Format(myTimeFormat))
	fwd, err := os.Create("../data/q85.dictt.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fwd.Close()

	enc = gob.NewEncoder(fwd)
	err = enc.Encode(dictt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\tfinished:", time.Now().Format(myTimeFormat))

	fmt.Println("save dictc", time.Now().Format(myTimeFormat))
	fwdc, err := os.Create("../data/q85.dictc.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fwd.Close()

	enc = gob.NewEncoder(fwdc)
	err = enc.Encode(dictc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\tfinished:", time.Now().Format(myTimeFormat))
}
