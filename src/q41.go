package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
	"strconv"
)

type Morph struct {
	surface, base, pos, pos1 string
}

type Chunk struct {
	morphs []Morph
	dst int
	srcs []int
}

func main()  {
	f, err := os.Open("../data/neko.txt.cabocha")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	
	r := bufio.NewReader(f)	
	
	sents := make([][]Chunk, 0)
	sent := make([]Chunk, 0)
	i := 0

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		if string(l) != "EOS" {
			if string(l[0]) == "*" {
				// split the read line by "D" and space to get the dst.
				tmp := strings.Split(string(l), "D")
				tmp = strings.Split(tmp[0], " ")
				dst, err := strconv.Atoi(tmp[len(tmp) - 1])
				if err != nil {
					panic(err)
				}
				
				c := Chunk {
					dst: dst,
				}
				
				sent = append(sent, c)
				i++
			} else {
				// split the read line by tab and camma, then store to sent[i-1].morphs as a list, where sent[i-1] is same as Chunk c.
				tmp := strings.Split(string(l), "\t")
				tmp = append(tmp[:1], strings.Split(tmp[1], ",")...)
				m := Morph {
					surface: tmp[0],
					base: tmp[7],
					pos: tmp[1],
					pos1: tmp[2],
				}
				sent[i - 1].morphs = append(sent[i - 1].morphs, m)
			}
		} else {
			// when read line is "EOS", then we process the initialization and gain the next sentence.
			if len(sent) > 0 { 
				for i, v := range sent {
					if v.dst != -1 {
						sent[v.dst].srcs = append(sent[v.dst].srcs, i)
					}
				}
				sents = append(sents, sent)
				sent = make([]Chunk, 0)
			}
			i = 0
			
		}
	}
	
	// print the 8th sentence.
	n := 7 // the sentence number we wanna print.
	for _, chunk := range sents[n] {
		for _, ms := range chunk.morphs {
			fmt.Print(ms.surface, " ")
		}
		if chunk.dst != -1 {
			fmt.Print("---> ")
			for _, md := range sents[n][chunk.dst].morphs {
				fmt.Print(md.surface, " ")
			}
		}
		fmt.Println("")
	}

}
