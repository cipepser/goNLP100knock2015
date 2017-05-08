package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Morph struct {
	surface, base, pos, pos1 string
}

type Chunk struct {
	morphs []Morph
	dst    int
	srcs   []int
}

// The function gotoNextChunk print the current morph and call itself recursively,
// it repeats until the destination will be -1.
func gotoNextChunk(sent []Chunk, now int) {
	chunk := sent[now]
	var morph string
	for _, m := range chunk.morphs {
		if m.pos != "記号" {
			morph += m.surface
		}
	}

	if chunk.dst != -1 {
		fmt.Print(morph, " -> ")
		gotoNextChunk(sent, chunk.dst)
	} else {
		fmt.Println(morph)
	}
}

func main() {
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
				dst, err := strconv.Atoi(tmp[len(tmp)-1])
				if err != nil {
					panic(err)
				}

				c := Chunk{
					dst: dst,
				}

				sent = append(sent, c)
				i++
			} else {
				// split the read line by tab and camma, then store to sent[i-1].morphs as a list, where sent[i-1] is same as Chunk c.
				tmp := strings.Split(string(l), "\t")
				tmp = append(tmp[:1], strings.Split(tmp[1], ",")...)
				m := Morph{
					surface: tmp[0],
					base:    tmp[7],
					pos:     tmp[1],
					pos1:    tmp[2],
				}
				sent[i-1].morphs = append(sent[i-1].morphs, m)
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

	// path from a noun to the root.
	for _, sent := range sents {
		for _, chunk := range sent {
			var morph string
			var flg bool

			// confirm if the chunk contains a noun.
			for _, m := range chunk.morphs {
				if m.pos != "記号" {
					morph += m.surface
				}
				if m.pos == "名詞" {
					flg = true
				}
			}

			if flg {
				if chunk.dst != -1 {
					fmt.Print(morph, " -> ")
					gotoNextChunk(sent, chunk.dst)
				} else {
					fmt.Println(morph)
				}
			}
		}
	}
}
