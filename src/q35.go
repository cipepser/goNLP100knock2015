package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
)

func main()  {
	f, err := os.Open("../data/neko.txt.mecab")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	
	r := bufio.NewReader(f)	
	
	sents := make([][]map[string]string, 0)
	sent := make([]map[string]string, 0)

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		
		// store morpheme which is not "EOS" into maps
		if string(b) != "EOS" {
			// split by tab and comma
			tmp := strings.Split(string(b), "\t")
			m := append(tmp[:1], strings.Split(tmp[1], ",")...)
			
			morpheme := make(map[string]string)
			
		 	morpheme["surface"] = m[0]
			morpheme["base"]    = m[7]
			morpheme["pos"]     = m[1]
			morpheme["pos1"]    = m[2]
			
			sent = append(sent, morpheme)
		} else { // if we find "EOS", store sentence to sentences and initialize the sent
			if len(sent) > 0 { // for appearing "EOS" continuously
				sents = append(sents, sent)
				sent = make([]map[string]string, 0)
			}
		}
		
	}
	
	res := make([]string, 0)
	for _, sent := range sents {
		i := 0
		for i < len(sent) - 1 {
			if sent[i]["pos"] == "名詞" {
				j := 1
				tmp := sent[i]["surface"]
				for sent[i + j]["pos"] == "名詞" {
					tmp += sent[i + j]["surface"]
					j++
					if i + j > len(sent) - 1 {
						break
					}
				}
				i += j
				res = append(res, tmp)
			} else {
				i++
			}
		}
	}
	
	fmt.Println(res[:3], res[len(res) - 3:])
}
