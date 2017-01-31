package main

import (
	"fmt"
	"sort"
	// "reflect"
)

func makeCharngram(s []rune, n int) runeSlice {
	t := make(runeSlice, len(s) - n + 1)
	
	for i := 0; i < len(s) -n + 1; i++ {
		t[i] = s[i:i+n]
	}
	
	return t
}

// convert runeSlice to []string
func String(R runeSlice) (s []string) {
	for _, r := range R {
		s = append(s, string(r))	
	}
	
	return
}

type runeSlice [][]rune

func (p runeSlice) Len() int {
	return len(p)
}
func (p runeSlice) Less(i, j int) bool {
	var isLess bool
	
	for k := 0; k < len(p[i]) && k < len(p[j]); k++ {
		if p[i][k] < p[j][k] {
			isLess = true
			break
		}
		if p[i][k] > p[j][k] {
			isLess = false
			break
		}		
	}
	return isLess
}

func (p runeSlice) Swap(i, j int) {
	 p[i], p[j] = p[j], p[i]
}

func main() {
	s1 := "paraparaparadise"
	
	X := makeCharngram([]rune(s1), 2)
	
	fmt.Println(X)
	sort.Sort(X)
	fmt.Println(X)
}