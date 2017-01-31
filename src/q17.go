package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
	"sort"
	"reflect"
)

const (
	NULL rune = 0
)

type runeSlice [][]rune

func (p runeSlice) Len() int {
	return len(p)
}

func (p runeSlice) Less(i, j int) bool {
	var isLess bool
	
	for k := 0; k < len(p[i]) && k < len(p[j]); k++ {
		if p[i][k] < p[j][k] {
			isLess = true
			break
		}
		if p[i][k] > p[j][k] {
			isLess = false
			break
		}		
	}
	return isLess
}

func (p runeSlice) Swap(i, j int) {
	 p[i], p[j] = p[j], p[i]
}

func String(R runeSlice) (s []string) {
	for _, r := range R {
		s = append(s, string(r))	
	}
	
	return
}

func Uniq(R runeSlice) runeSlice {
	i := 0
	for {
		if reflect.DeepEqual(R[i], R[i + 1]) {
			copy(R[i:], R[i + 1:])
			R[len(R) - 1] = nil
			R = R[:len(R) - 1]
		} else {
			i++
		}

		if i == len(R) - 1 {
			break
		}
	}

	return R
}

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
	
	// col列目を抜き出す
	var count int = 1
	var word []rune
	var words runeSlice
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
			words = append(words, word)
			count = 1
			word = nil
		}
	}
	
	// sort
	sort.Sort(words)
	
	// uniq
	words = Uniq(words)
	
	// 結果表示
	fmt.Println(String(words))
}