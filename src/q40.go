package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
)

type Morph struct {
	surface, base, pos, pos1 string
}

func main()  {
	f, err := os.Open("../data/neko.txt.cabocha")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	
	r := bufio.NewReader(f)	
	
	sent := make([]Morph, 0)
	sents := make([][]Morph, 0)

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		if string(l[0]) != "*" {
			if string(l) != "EOS" {
				// split the read line by tab and camma, then stored as class Morph.
				tmp := strings.Split(string(l), "\t")
				tmp = append(tmp[:1], strings.Split(tmp[1], ",")...)
				m := Morph {
					surface: tmp[0],
					base: tmp[7],
					pos: tmp[1],
					pos1: tmp[2],
				}
				
				sent = append(sent, m)
			} else { // when read line is "EOS", we gain to next sentence.
				if len(sent) > 0 { 
					sents = append(sents, sent)
					sent = make([]Morph, 0)
				}
			}
		}
	}
	
	// print the 3rd sentence.
	fmt.Println(sents[2])
}
