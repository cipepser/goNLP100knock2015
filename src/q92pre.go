package main

import (
	"bufio"
	"encoding/gob"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

var (
	dim = 300
)

func main() {
	frp, err := os.Open("../data/q85.proj.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer frp.Close()

	dec := gob.NewDecoder(frp)

	var proj mat.Dense
	err = dec.Decode(&proj)
	if err != nil {
		log.Fatal(err)
	}

	frd, err := os.Open("../data/q85.dictt.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer frd.Close()

	dec = gob.NewDecoder(frd)

	dict := make(map[string]int)
	err = dec.Decode(&dict)
	if err != nil {
		log.Fatal(err)
	}

	fw, err := os.Create("../data/q85_model.txt")
	defer fw.Close()
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriterSize(fw, 4096)
	defer w.Flush()

	w.Write([]byte(strconv.Itoa(len(dict)) + " " + strconv.Itoa(dim) + "\n"))

	for word, idx := range dict {
		w.Write([]byte(word + " "))
		vec := proj.RowView(idx)
		for i := 0; i < dim; i++ {
			w.Write([]byte(strconv.FormatFloat(vec.At(i, 0), 'f', 6, 64) + " "))
		}
		w.Write([]byte("\n"))
	}
}
