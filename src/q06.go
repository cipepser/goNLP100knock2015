package main

import (
	"fmt"
	"reflect"
)

func makeCharngram(s []rune, n int) [][]rune {
	t := make([][]rune, len(s) - n + 1)
	
	for i := 0; i < len(s) -n + 1; i++ {
		t[i] = s[i:i+n]
	}
	
	return t
}

func Exists(set [][]rune, word []rune) bool {
	var b bool
	for _, v := range set {
		if reflect.DeepEqual(v, word) {
			b = true
			break
		}
	}
	return b
}

// 和集合
func Union(X, Y [][]rune) [][]rune {
	for _, y := range Y {
		if !Exists(X, y) {
			X = append(X, y)
		}
	}
	
	return X
}

// 積集合
func Intersection(X, Y [][]rune) [][]rune {
	var Z [][]rune
	for _, y := range Y {
		if Exists(X, y) {
			Z = append(Z, y)
		}
	}
	
	return Z
}

// 差集合(X - Y)
func Difference(X, Y [][]rune) [][]rune {
	var Z [][]rune
	for _, x := range X {
		if !Exists(Y, x) {
			Z = append(Z, x)
		}
	}
		
	return Z
}

// convert [][]rune to []string
func String(R [][]rune) (s []string) {
	for _, r := range R {
		s = append(s, string(r))	
	}
	
	return
}


func main()  {
	s1 := "paraparaparadise"
	s2 := "paragraph"
	s := "se"
	
	X := makeCharngram([]rune(s1), 2)
	Y := makeCharngram([]rune(s2), 2)
	
	fmt.Println(String(X))
	fmt.Println(String(Y))
	fmt.Println(String(Intersection(X, Y)))
	fmt.Println(String(Union(X, Y)))
	fmt.Println(String(Difference(X, Y)))
	
	Z := makeCharngram([]rune(s), 2)
	fmt.Println(Exists(X, Z[0]))
	fmt.Println(Exists(Y, Z[0]))
}