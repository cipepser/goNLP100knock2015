package main

import (
	"fmt"
	"strings"
)

func main()  {
	var s string = "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, ",", "", -1)

	for _, n := range strings.Split(s, " ") {
		fmt.Println(len(n))
	}
}