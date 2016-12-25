package main

import (
	"fmt"
)

func Reverse(s string) string {
	t := make([]rune, len(s))
	for i, c := range s {
		t[len(s) - i- 1] = c
	} 
	return string(t)
}

func main()  {
	var s string = "stressed"
	
	fmt.Println(Reverse(s))
}