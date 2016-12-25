package main

import (
	"fmt"
)

func main()  {
	var s string = "パタトクカシーー"
	
	t := []rune(s)			
	fmt.Println(string(t[1]) + string(t[3]) + string(t[5]) + string(t[7]))	
}