package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"io"
)

func main()  {
	var count int

	if len(os.Args) != 2 {
		log.Fatal("Please input filename as a first standard input.")
	} else {
		fmt.Printf(">> read file: %s\n", os.Args[1])
		fp, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()

		reader := bufio.NewReaderSize(fp, 4096)

		for {
			_, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			count++
		}
		
		fmt.Println(count)
	}
}