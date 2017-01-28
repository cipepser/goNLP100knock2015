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

	dir, _ := filepath.Split(os.Args[1])
	
	nFile := row / N + 1
	if row % N == 0 {
		nFile--
	}
	
	wfp := make([]*os.File, nFile)

	for i := 0; i < nFile; i++ {
		filename := dir + "div" + strconv.Itoa(i) + ".txt"
		wfp[i], err = os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer wfp[i].Close()
	}
	
	var i, count int
	for {
		p, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		count++

		wfp[i].Write(p)
		wfp[i].Write([]byte(string('\n')))

		if count % N == 0 {
			i++
		}
	}
}