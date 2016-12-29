package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
)

func FisherYatesShuffle (s string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano() ))
	t := []rune(s)
	
	var j int
	for i := 0; i < len(t) - 1; i++ {
		j = r.Intn(len(t) - 1 - i)
		t[len(t) - 1 - i], t[j] = t[j], t[len(t) - 1 - i] 
	}
	
	return string(t)
}


func main()  {
	var s string = "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ." 

	var u string
	var v_tmp []rune
	for _, v := range strings.Split(s, " ") {
		if len(v) > 4 {
			v_tmp = []rune(v)
			v = string(v_tmp[0])
			v += FisherYatesShuffle(string(v_tmp[1:len(v_tmp)-1]))
			v += string(v_tmp[len(v_tmp)-1])
		}
		u += v + " "
	}

	u = u[0:len(u)-1] // remove last blank
	
	fmt.Println(u)
}