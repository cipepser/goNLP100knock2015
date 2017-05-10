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

func getIdcesToRoot(sent []Chunk, now int, pos []int) []int {
	chunk := sent[now]
	pos = append(pos, now)

	if chunk.dst != -1 {
		return getIdcesToRoot(sent, chunk.dst, pos)
	} else {
		return pos
	}
}

func getFirstIntersection(as, bs []int) int {
	for i, a := range as {
		for j, b := range bs {
			if a == b {
				if i == 0 || j == 0 { // This is same as directPath.
					return -1
				} else {
					return a
				}
			}
		}
	}
	return -1
}

func indirectNounPairs(sent []Chunk) {
	var pathes [][]int

	for i, chunk := range sent {
		for _, m := range chunk.morphs {
			if m.pos == "名詞" && chunk.dst != -1 {
				pathes = append(pathes, getIdcesToRoot(sent, chunk.dst, []int{i}))
				break
			}
		}
	}

	for i := 0; i < len(pathes)-1; i++ {
		for j := i + 1; j < len(pathes); j++ {
			k := getFirstIntersection(pathes[i], pathes[j])
			if k > 0 {
				for l := 0; l < len(pathes[i])-1; l++ {
					var morph string
					var flg bool
					for _, m := range sent[pathes[i][l]].morphs {
						if m.pos != "記号" {
							if m.pos == "名詞" && l == 0 && !flg {
								morph += "X"
								flg = true
							} else {
								morph += m.surface
							}
						}
					}

					fmt.Print(morph)

					if l < len(pathes[i])-2 {
						fmt.Print(" -> ")
					}
				}

				fmt.Print(" | ")

				for l := 0; l < len(pathes[j])-1; l++ {
					var morph string
					var flg bool
					for _, m := range sent[pathes[j][l]].morphs {
						if m.pos != "記号" {
							if m.pos == "名詞" && l == 0 && !flg {
								morph += "Y"
								flg = true
							} else {
								morph += m.surface
							}
						}
					}

					fmt.Print(morph)

					if l < len(pathes[j])-2 {
						fmt.Print(" -> ")
					}
				}

				fmt.Print(" | ")

				for _, m := range sent[k].morphs {
					if m.pos != "記号" {
						fmt.Print(m.surface)
					}
				}
				fmt.Println("")
			}
		}
	}
}

func directNounPairs(sent []Chunk) {
	for _, chunk := range sent {
		var morph string
		var flg bool
		// confirm if the chunk contains a noun.
		for _, m := range chunk.morphs {
			if m.pos != "記号" {
				if m.pos == "名詞" && !flg {
					morph += "X"
					flg = true
				} else {
					morph += m.surface
				}
			}
		}

		if flg {
			if chunk.dst != -1 {
				for _, end := range getIdxOfNoun(sent, chunk.dst, []int{}) {
					fmt.Print(morph, " -> ")
					gotoNextChunk(sent, chunk.dst, end)
				}
			}
		}
	}
}

func gotoNextChunk(sent []Chunk, now, end int) {
	chunk := sent[now]
	var morph string
	for _, m := range chunk.morphs {
		if m.pos != "記号" {
			morph += m.surface
		}
	}

	if now == end {
		fmt.Println("Y")
		return
	}

	if chunk.dst != -1 {
		fmt.Print(morph, " -> ")
		gotoNextChunk(sent, chunk.dst, end)
	}
}

func getIdxOfNoun(sent []Chunk, now int, pos []int) []int {
	chunk := sent[now]

	for _, m := range chunk.morphs {
		if m.pos == "名詞" {
			pos = append(pos, now)
		}
	}

	if chunk.dst != -1 {
		return getIdxOfNoun(sent, chunk.dst, pos)
	} else {
		return pos
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

	for _, sent := range sents {
		indirectNounPairs(sent)
		directNounPairs(sent)
	}
}
