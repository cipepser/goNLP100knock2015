package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"log"
	"strconv"
)

func main() {
	rfp, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer rfp.Close()

	N, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReaderSize(rfp, 4096)
	var count int
	
	for count < N {
		p, _, err := reader.ReadLine()
		count++
		if err == io.EOF {
			break
		}
		fmt.Println(string(p))
	}
}