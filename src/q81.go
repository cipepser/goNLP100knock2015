package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func connectCountryName(s string) string {
	countries := []string{"American Samoa", "Antigua and Barbuda", "Ashmore and Cartier Islands", "Bahamas, The",
		"Bassas da India", "Bosnia and Herzegovina", "Bouvet Island", "British Indian Ocean Territory", "British Virgin Islands",
		"Burkina Faso", "Cape Verde", "Cayman Islands", "Central African Republic", "Christmas Island", "Clipperton Island",
		"Cocos (Keeling) Islands", "Congo, Democratic Republic of the", "Congo, Republic of the", "Cook Islands", "Coral Sea Islands",
		"Costa Rica", "Cote d'Ivoire", "Czech Republic", "Dominican Republic", "El Salvador", "Equatorial Guinea", "Europa Island",
		"Falkland Islands (Islas Malvinas)", "Faroe Islands", "French Guiana", "French Polynesia", "French Southern and Antarctic Lands",
		"Gambia, The", "Gaza Strip", "Glorioso Islands", "Heard Island and McDonald Islands", "Holy See (Vatican City)", "Hong Kong",
		"Isle of Man", "Jan Mayen", "Juan de Nova Island", "Korea, North", "Korea, South", "Marshall Islands", "Micronesia, Federated States of",
		"Navassa Island", "Netherlands Antilles", "New Caledonia", "New Zealand", "Norfolk Island", "Northern Mariana Islands",
		"Papua New Guinea", "Paracel Islands", "Pitcairn Islands", "Puerto Rico", "Saint Helena", "Saint Kitts and Nevis", "Saint Lucia",
		"Saint Pierre and Miquelon", "Saint Vincent and the Grenadines", "San Marino", "Sao Tome and Principe", "Saudi Arabia",
		"Serbia and Montenegro", "Sierra Leone", "Solomon Islands", "South Africa", "South Georgia and the South Sandwich Islands",
		"Spratly Islands", "Sri Lanka", "Trinidad and Tobago", "Tromelin Island", "Turks and Caicos Islands", "United Arab Emirates",
		"United Kingdom", "United States", "Virgin Islands", "Wake Island", "Wallis and Futuna", "West Bank", "Western Sahara"}

	for _, c := range countries {
		s = strings.Replace(s, c,
			strings.Replace(c, " ", "_", -1),
			-1)
	}

	return s
}

func main() {
	f, err := os.Open("../data/q80.out.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	corpus := [][]string{}
	r := bufio.NewReaderSize(f, 4096)
	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		s := connectCountryName(string(l))

		c := []string{}
		tokens := strings.Split(s, " ")
		for _, t := range tokens {
			c = append(c, t)
		}
		if len(c) > 0 {
			corpus = append(corpus, c)
		}
	}

	// write the result to txt file
	fw, err := os.Create("../data/q81.out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fw.Close()

	for _, ts := range corpus {
		for i, t := range ts {
			fw.Write([]byte(t))
			if i != len(ts)-1 {
				fw.Write([]byte(string(' ')))
			}
		}
		fw.Write([]byte("\n"))
	}
}
