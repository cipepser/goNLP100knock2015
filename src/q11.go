package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	// "io"
)

const (
	NULL rune = 0
)

func main()  {
	if len(os.Args) != 2 {
		log.Fatal("Please input filename as a first standard input.")
	} else {
		fp, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()

		reader := bufio.NewReaderSize(fp, 4096)
		
		for {
			p, _, _ := reader.ReadRune()
			
			if p == rune('\t') {
				p = rune(' ')
			}

			if p == NULL {
				break
			}
			fmt.Printf(string(p))
		}
		
	}
}