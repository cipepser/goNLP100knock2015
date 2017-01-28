package main

import (
	"os"
	"bufio"
	"log"
	"strconv"
	"path/filepath"
	"io"
)

func CountRow(filename string) (count int) {
	fp, err := os.Open(filename)
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
	
	return
}

func main() {
	// read file
	rfp, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer rfp.Close()

	reader := bufio.NewReaderSize(rfp, 4096)
	row := CountRow(os.Args[1])

	N, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	// write file
	dir, _ := filepath.Split(os.Args[1])
	wfp := make([]*os.File, N)

	for i := 0; i < N; i++ {
		filename := dir + "div" + strconv.Itoa(i) + ".txt"
		wfp[i], err = os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer wfp[i].Close()
	}
	
	for count := 0; count < row; count++ {
		p, _, _ := reader.ReadLine()
		
		wfp[count / (row / N + 1)].Write(p)
		wfp[count / (row / N + 1)].Write([]byte(string('\n')))
	}
}