package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
	"sort"
)

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}
func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}
func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

// sortedKeys returns the slice []string which is sorted by the count in map.
// Eventually, we should hold the sorted slice []string, it's enough to sort.
// func (sm *sortedMap) sortedKeys(m map[string]int) []string {
func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	
	sort.Sort(sm)
	
	return sm.s
}

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
	
	// count the number of same base of morpheme
	freq := make(map[string]int)
	for _, sent := range sents {
		for _, m := range sent {
			freq[m["base"]]++
		}
	}
	
	// For printing all the bases of morpheme
	// for _, r := range sortedKeys(freq) {
	// 	fmt.Println(r, freq[r])
	// }
	
	res := sortedKeys(freq)	
	for _, v := range res[:10] {
		fmt.Println(v, freq[v])
	}
	
}
