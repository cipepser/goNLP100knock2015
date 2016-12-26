package main

import (
	"fmt"
	"strconv"
)

func Template(x int, y string, z float64) string {
	var s = strconv.Itoa(x) + "時の" + y + "は" + strconv.FormatFloat(z, 'f', -1, 64)
	return s
}

func main()  {
	var x int = 12
	var y string = "気温"
	var z float64 = 22.4
	
	fmt.Println(Template(x, y, z))
}