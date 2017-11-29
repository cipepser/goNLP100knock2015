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

	"gonum.org/v1/gonum/mat"

	"github.com/james-bowman/sparse"
	"gonum.org/v1/gonum/stat"
)

const (
	myTimeFormat = "2006/01/02 15:04:05 JST"
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
	fmt.Println("start")
	fmt.Println(time.Now().Format(myTimeFormat))

	fmt.Println("read file")
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
	fmt.Println(time.Now().Format(myTimeFormat))

	fmt.Println("calc PPMI")
	X := make(map[key]float64)
	for k := range m {
		ppmi := calcPPMI(len(m), m[k], Nt[k.t], Nc[k.t])
		if ppmi > 0 {
			X[k] = ppmi
		}
	}
	fmt.Println(time.Now().Format(myTimeFormat))

	fmt.Println("the number of the word: ", len(Nt))
	fmt.Println("the number of contex word: ", len(Nc))

	// transform hashmap to matrix
	fmt.Println("re-order idxt")
	idxt := make(map[int]string)
	dictt := make(map[string]int)
	i := 0
	for k := range Nt {
		idxt[i] = k
		dictt[k] = i
		i++
	}
	fmt.Println(time.Now().Format(myTimeFormat))

	fmt.Println("re-order idxc")
	idxc := make(map[int]string)
	dictc := make(map[string]int)
	j := 0
	for k := range Nc {
		idxc[j] = k
		dictc[k] = j
		j++
	}
	fmt.Println(time.Now().Format(myTimeFormat))

	fmt.Println("store data as a CSR")
	// TODO: 高速化

	fmt.Println(len(X))
	for k, v := range X {
		fmt.Println("(", dictt[k.t], ",", dictc[k.c], "): ", v)
	}

	// ja = make([]int, len(X))
	data := []float64{}
	// ia := []int{}
	ia := make([]int, len(Nt)+1)
	ja := []int{}
	cnt := 0
	for i := 0; i < len(Nt); i++ {
		for j := 0; j < len(Nc); j++ {
			k := key{
				t: idxt[i],
				c: idxc[j],
			}
			if X[k] > 0 {
				data = append(data, X[k])
				ja = append(ja, j)
				cnt++
			}
		}
		ia[i+1] = cnt

	}
	fmt.Println(time.Now().Format(myTimeFormat))

	fmt.Println("new CSR")
	// TODO: NewDenseされてmakesliceのout of rangeになる
	y := sparse.NewCSR(len(Nt), len(Nc), ia, ja, data)
	fmt.Println(time.Now().Format(myTimeFormat))

	// PCA
	fmt.Println("PCA start...")
	var pc stat.PC
	ok := pc.PrincipalComponents(y, nil)
	if !ok {
		log.Fatal("PCA fails")
	}

	fmt.Println(time.Now().Format(myTimeFormat))
	fmt.Println("PCA finshed!!")

	k := 300
	var proj mat.Dense
	proj.Mul(y, pc.VectorsTo(nil).Slice(0, len(Nc), 0, k))

	fmt.Println("save proj...")
	// save the result
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
	fmt.Println(time.Now().Format(myTimeFormat))

	fmt.Println("save dictt...")
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
	fmt.Println(time.Now().Format(myTimeFormat))

	fmt.Println("save dictc...")
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
	fmt.Println(time.Now().Format(myTimeFormat))

}
