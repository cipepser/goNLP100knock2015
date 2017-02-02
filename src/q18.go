package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
	"sort"
)

const (
	NULL rune = 0
)

// define Record as each row
type Record struct {
	Pref []rune
	City []rune
	Temperature float64
	Date []rune
}

type Table []Record

// sort for Table
func (t Table) Len() int {
	return len(t)
}
func (t Table) Swap(i, j int) {
	 t[i], t[j] = t[j], t[i]
}

type ByPref struct {
	Table
}
func (b ByPref) Less(i, j int) bool {
	return isLess(b.Table[i].Pref, b.Table[j].Pref)
}

type ByCity struct {
	Table
}
func (b ByCity) Less(i, j int) bool {
	return isLess(b.Table[i].City, b.Table[j].City)
}

type ByTemperature struct {
	Table
}
func (b ByTemperature) Less(i, j int) bool {
	return b.Table[i].Temperature < b.Table[j].Temperature
}

type ByDate struct {
	Table
}
func (b ByDate) Less(i, j int) bool {
	return isLess(b.Table[i].Date, b.Table[j].Date)
}

func isLess(r1, r2 []rune) bool {
	var isLess, flgSwap bool
	
	if len(r1) > len(r2) {
		r1, r2 = r2, r1
		flgSwap = true
	}
	
	for i := 0; i < len(r1); i++ {
		if r1[i] < r2[i] {
			isLess = true
			break
		}
		if r1[i] > r2[i] {
			isLess = false
			break
		}
		
		// last loop
		if i == len(r1) - 1 {
			isLess = true
		}
	}
	
	if flgSwap {
		return !isLess
	} else {
		return isLess	
	}	
}

// print rune slice
func String(R []rune) []string {
	s := make([]string, len(R))
	for _, r := range R {
		s = append(s, string(r))	
	}
	
	return s
}

func main() {
	// file pointor to read file
	rfp, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer rfp.Close()

	reader := bufio.NewReaderSize(rfp, 4096)

	// sort by the col
	col, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	
	// read file
	var i int
	var table Table
	var record Record
	var temperature []rune

	for {
		
		p, _, _ := reader.ReadRune()
		if p == NULL {
			break
		}
		if p == rune('\t') {
			i++
			if i == 3 {
				record.Temperature, _ =  strconv.ParseFloat(string(temperature), 64)
			}
		} else {
			switch i {
			case 0:
				record.Pref = append(record.Pref, p)
			case 1:
				record.City = append(record.City, p)
			case 2:
				temperature = append(temperature, p)
			case 3:
				record.Date = append(record.Date, p)
			}
		}
		
		if p == rune('\n') {
			table = append(table, record)

			record = Record{}
			temperature = nil
			i = 0
		}
	}
	
	// sort
	switch col {
	case 1:
		sort.Sort(ByPref{table})
	case 2:
		sort.Sort(ByCity{table})
	case 3:
		sort.Sort(ByTemperature{table})
	case 4:
		sort.Sort(ByDate{table})
	}
	
	// result
	for _, r := range table {
		fmt.Printf("%s\t", string(r.Pref))
		fmt.Printf("%s\t", string(r.City))
		fmt.Printf("%.1f\t", r.Temperature)
		fmt.Printf("%s", string(r.Date))
	}
}