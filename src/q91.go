package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	fr, err := os.Open("../data/questions-words.txt")
	defer fr.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReaderSize(fr, 4096)

	fw, err := os.Create("../data/q91.out.txt")
	defer fw.Close()
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriterSize(fw, 4096)
	defer w.Flush()

	flg := false

	delimiter := ": "
	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if strings.Contains(string(l), delimiter) {
			if flg {
				break
			}
			if strings.Contains(string(l), "family") {
				flg = true
			}
		}

		if flg {
			w.Write(l)
			w.Write([]byte("\n"))
		}
	}
}
