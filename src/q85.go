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

	"gonum.org/v1/gonum/mat"

	"github.com/james-bowman/sparse"
	"gonum.org/v1/gonum/stat"
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

	X := make(map[key]float64)
	for k := range m {
		X[k] = calcPPMI(len(m), m[k], Nt[k.t], Nc[k.t])
	}

	fmt.Println("the number of the word: ", len(Nt))
	fmt.Println("the number of contex word: ", len(Nc))

	// transform hashmap to matrix
	idxt := make(map[int]string)
	dict := make(map[string]int)
	i := 0
	for k := range Nt {
		idxt[i] = k
		dict[k] = i
		i++
	}

	idxc := make(map[int]string)
	j := 0
	for k := range Nc {
		idxc[j] = k
		j++
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

	// ia[len(ia)-1] = len(data)

	// fmt.Println(len(ia))
	// fmt.Println(len(ja))
	//
	// fmt.Println(ia)
	// fmt.Println(ja)

	// data := []float64{3, 1, 1, 2, 1, 2, 1, 4}
	//
	// ia := []int{0, 2, 4, 6, len(data)}
	// ja := []int{0, 1, 0, 3, 2, 3, 0, 3}

	y := sparse.NewCSR(len(Nt), len(Nc), ia, ja, data)
	// _ = y
	//
	//
	//

	// fmt.Println(y.Dims())

	// y := mat.NewDense(len(Nt), len(Nc), nil)
	// for i := 0; i < len(Nt)-1; i++ {
	// 	for j := 0; j < len(Nc)-1; j++ {
	// 		k := key{
	// 			t: idxt[i],
	// 			c: idxc[j],
	// 		}
	// 		y.Set(i, j, X[k])
	// 	}
	// }
	//

	// PCA
	fmt.Println("PCA start...")
	var pc stat.PC
	ok := pc.PrincipalComponents(y, nil)
	if !ok {
		log.Fatal("PCA fails")
	}

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

	fmt.Println("save dict...")
	fwd, err := os.Create("../data/q85.dict.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fwd.Close()

	enc = gob.NewEncoder(fwd)

	err = enc.Encode(dict)
	if err != nil {
		log.Fatal(err)
	}

}
