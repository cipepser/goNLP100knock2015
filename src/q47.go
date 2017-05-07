package main

import (
	"bufio"
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

	// write the result to q47.out.txt
	fo, err := os.Create("../data/q47.out.txt")
	defer fo.Close()
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(fo)

	// connect the nodes.
	for _, sent := range sents {
		for _, chunk := range sent {
			for _, m := range chunk.morphs {
				if len(chunk.srcs) == 0 {
					continue
				}

				// write predicate.
				if m.pos == "動詞" {
					flg := false
					var predicate string
					idx_wo := -1
					cnt := 0
					for _, s := range chunk.srcs {
						idx_wo = s
						for i, c := range sent[s].morphs {
							if c.pos == "助詞" {
								if c.surface == "を" {
									if sent[s].morphs[i-1].pos == "名詞" && sent[s].morphs[i-1].pos1 == "サ変接続" {
										flg = true
										predicate = sent[s].morphs[i-1].base + c.base + m.base
									}
								}
								cnt++
							}
						}
					}

					if !flg {
						continue
					} else if cnt == 1 {
						continue
					}

					_, err := w.Write([]byte(predicate))
					if err != nil {
						panic(err)
					}

					// write postpositional particle.
					_, err = w.Write([]byte("\t"))
					if err != nil {
						panic(err)
					}

					cnt = 0
					for _, s := range chunk.srcs {
						idx_last_post := -1 // if we find two or more postpositional particle in a chunk, we should write the last one.
						for i, c := range sent[s].morphs {
							if c.pos == "助詞" && s != idx_wo {
								idx_last_post = i
							}
						}
						if cnt > 0 {
							_, err := w.Write([]byte(" "))
							if err != nil {
								panic(err)
							}
						}
						cnt++
						if idx_last_post != -1 {
							_, err := w.Write([]byte(sent[s].morphs[idx_last_post].surface))
							if err != nil {
								panic(err)
							}
						}
					}

					// write morphs which follow the case pattern.
					_, err = w.Write([]byte("\t"))
					if err != nil {
						panic(err)
					}

					cnt = 0
					for _, s := range chunk.srcs {
						var morph string
						idx_last_post := -1
						for i, c := range sent[s].morphs {
							if c.pos != "記号" {
								morph += c.surface
							}
							if c.pos == "助詞" && s != idx_wo {
								idx_last_post = i
							}
							if cnt > 0 {
								_, err := w.Write([]byte(" "))
								if err != nil {
									panic(err)
								}
							}
							cnt++
						}
						if idx_last_post != -1 {
							_, err := w.Write([]byte(morph))
							if err != nil {
								panic(err)
							}
						}
					}

					_, err = w.Write([]byte("\n"))
					if err != nil {
						panic(err)
					}
					if flg {
						break
					}
				}
			}
		}
		err := w.Flush() // write buffered data to the output file.
		if err != nil {
			panic(err)
		}
	}
}
