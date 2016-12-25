package main

import (
	"fmt"
)

func AltnativeConnect(s1, s2 string) string {
	// len(s1) = len(s2)
	t := make([]rune, len(s1) + len(s2))
	
	r1 := []rune(s1)
	r2 := []rune(s2)
	
	for i := 0; i < len(r2); i++ {
		t[2 * i]     = r1[i]
		t[2 * i + 1] = r2[i]
	}
	
	return string(t)
}


func main()  {
	var s1, s2 string = "パトカー", "タクシー"
	
	fmt.Println(AltnativeConnect(s1, s2))
}