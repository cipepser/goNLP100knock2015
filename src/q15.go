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
	output := make([][]byte, N)
	
	for {
		p, _, err := reader.ReadLine()
		count++
		if err == io.EOF {
			break
		}

		output[(count-1) % N] = p
	}
	
	if count < N {
		output = output[:count - 1]
	}
	
	i := (count - 1) % N
	
	output = append(output[i:], output[:i]...)

	for _, v := range output {
		fmt.Println(string(v))
	}
}