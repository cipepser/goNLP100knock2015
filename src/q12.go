package main

import (
	"os"
	"bufio"
	"log"
	"strconv"
	"path/filepath"
)

const (
	NULL rune = 0
)

func main() {
	// read file
	rfp, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer rfp.Close()

	reader := bufio.NewReaderSize(rfp, 4096)

	// 抜き出したい列
	col, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	// write file
	dir, _ := filepath.Split(os.Args[1])
	filename := dir + "col" + strconv.Itoa(col) + ".txt"
	wfp, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer wfp.Close()

	
	// output := []byte("aaaa")
	// wfp.Write(output)

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
		} else if count == col {
			word = append(word, p)
		}
		
		if p == rune('\n') {
			wfp.Write([]byte(string(word)))
			wfp.Write([]byte(string('\n')))
			count = 1
			word = nil
		}

	}
}