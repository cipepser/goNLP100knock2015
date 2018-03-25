package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"gonum.org/v1/gonum/mat"
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

	word := "United_States"
	fmt.Println(word, ":", proj.RowView(dict[word]))
}
