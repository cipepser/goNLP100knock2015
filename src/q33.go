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
		for _, m := range sent {
			if m["pos1"] == "サ変接続" {
				res = append(res, m["surface"])
			}
		}
	}
	
	fmt.Println(res[:3], res[len(res) - 3:])
}
