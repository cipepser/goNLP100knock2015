package main

import (
	"fmt"
	"strings"
)

func makeCharngram(s []rune, n int) [][]rune {
	t := make([][]rune, len(s) - n + 1)
	
	for i := 0; i < len(s) -n + 1; i++ {
		t[i] = s[i:i+n]
	}
	
	return t
}

func makeWordngram(s [][]rune, n int) [][][]rune {
	t := make([][][]rune, len(s) - n + 1)
	
	for i := 0; i < len(s) -n + 1; i++ {
		t[i] = s[i:i+n]
	}
	
	return t
}

func main()  {
	s := "I am an NLPer"
	
	fmt.Println("------文字bi-gram------")
	t := makeCharngram([]rune(s), 2)
	for _, v := range t {
		fmt.Println(string(v))
	}
	
	fmt.Println("------単語bi-gram------")
	words := strings.Split(s, " ")
	r_words := make([][]rune, len(words))

	for i, v := range words {
		r_words[i] = []rune(v)
	}
	wbg := makeWordngram(r_words, 2)
	
	for _, w := range wbg {
		for _, x := range w {
			fmt.Printf(string(x) + " ")
		} 
		fmt.Println("")
	}
	
}