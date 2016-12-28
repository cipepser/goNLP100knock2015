package main

import (
	"fmt"
)

func cipher(r rune) string {
	var u rune
	if r >= []rune("a")[0] && r <= []rune("z")[0] {
		u = 219 - r
	} else {
		u = r
	}
	
	return string(u)
}

func main()  {
	plain := "This is a pen."
	for _, s := range plain {
		fmt.Println(cipher(s))
	}
}