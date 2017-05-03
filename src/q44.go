package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"github.com/awalterschulze/gographviz"
)

type Morph struct {
	surface, base, pos, pos1 string
}

type Chunk struct {
	morphs []Morph
	dst    int
	srcs   []int
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

	// draw the directed graph. 
	n := 10 // the number of sentence you wanna print.
	// fmt.Println(sents[n])
	g := gographviz.NewGraph()
	if err := g.SetName("G"); err != nil {
		panic(err)
	}
	if err := g.SetDir(true); err != nil { 
		panic(err)
	}
	if err := g.AddAttr("G", "bgcolor", "\"#343434\""); err != nil { 
		panic(err)
	}
	
	// configuration for nodes
	nodeAttrs := make(map[string]string)
	nodeAttrs["colorscheme"] = "rdylgn11"
	nodeAttrs["style"] = "\"solid,filled\""
	nodeAttrs["fontcolor"] = "6"
	nodeAttrs["fontname"] = "\"Migu 1M\""
	nodeAttrs["color"] = "7"
	nodeAttrs["fillcolor"] = "11"
	nodeAttrs["shape"] = "doublecircle"
	
	// add the chunk as a node.
	for _, chunk := range sents[n] {
		var ws string
		for _, ms := range chunk.morphs {
			if ms.pos != "記号" {
				ws += ms.surface
			}
		}
		if err := g.AddNode("G", ws, nodeAttrs); err != nil {
			panic(err)
		}		
	}

	// configuration for edges
	edgeAttrs := make(map[string]string)
	edgeAttrs["color"] = "white"
	
	// connect the nodes.
	for _, chunk := range sents[n] {
		var ws, wd string

		for _, ms := range chunk.morphs {
			if ms.pos != "記号" {
				ws += ms.surface
			}
		}
		if chunk.dst != -1 {
			for _, md := range sents[n][chunk.dst].morphs {
				if md.pos != "記号" {
					wd += md.surface
				}
			}
			if err := g.AddEdge(ws, wd, true, edgeAttrs); err != nil {
				panic(err)
			}
		}
	}

	// output the dotfile.
	s := g.String()
	file, err := os.Create(`../data//q44_diGraph.dot`)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte(s))

}
