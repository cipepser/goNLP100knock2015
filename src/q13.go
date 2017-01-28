package main

import (
	"os"
	"bufio"
	"io"
	"log"
)

const (
	dir string = "../data/"
)

func main() {
	// read file
	rfp1, err := os.Open(dir + "col1.txt")
	defer rfp1.Close()
	rfp2, err := os.Open(dir + "col2.txt")
	defer rfp2.Close()

	reader1 := bufio.NewReaderSize(rfp1, 4096)
	reader2 := bufio.NewReaderSize(rfp2, 4096)

	// write file
	filename := dir + "col_merge.txt"
	wfp, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer wfp.Close()

	for {
		p1, _, err := reader1.ReadLine()
		if err == io.EOF {
			break
		}
		
		p2, _, err := reader2.ReadLine()
		if err == io.EOF {
			break
		}

		wfp.Write(p1)
		wfp.Write([]byte(string('\t')))
		wfp.Write(p2)
		wfp.Write([]byte(string('\n')))
	}
}