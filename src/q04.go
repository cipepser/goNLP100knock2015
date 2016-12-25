package main

import (
	"fmt"
	"strings"
)

func main()  {
	var s string = "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can." 
	s = strings.Replace(s, ".", "", -1)
	
	dict := make(map[string]int)
	for i, v := range strings.Split(s, " ") {
		switch i + 1 {
		case 1, 5, 6, 7, 8, 9, 15, 16, 19:
			dict[string([]rune(v)[0])] = i
		default:
			dict[string([]rune(v)[0]) + string([]rune(v)[1])] = i
		}
	}
	fmt.Println(dict)
}