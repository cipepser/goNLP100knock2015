package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
)

const (
	NULL rune = 0
)

func main()  {
	if len(os.Args) != 3 {
		log.Fatal("Please input filename as a first standard input and row number you want to extract as a second.")
	} else {
		fp, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()

		reader := bufio.NewReaderSize(fp, 4096)

		row, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		var count int = 1
		var word []rune		
		for {
			p, _, _ := reader.ReadRune()
			// ファイル終端ならループを抜ける
			if p == NULL {
				break
			}

			if p == rune('\t') {
				count++
			} else if count == row {
				word = append(word, p)
			}
			
			if p == rune('\n') {
				fmt.Printf("%v\n", string(word))
				count = 1
				word = nil
			}

		}
		
	}
}