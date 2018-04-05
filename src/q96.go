package main

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	dim       = 300
	countries = []string{"American_Samoa", "Antigua_and_Barbuda", "Ashmore_and_Cartier_Islands", "Bahamas,The",
		"Bassas_da_India", "Bosnia_and_Herzegovina", "Bouvet_Island", "British_Indian_Ocean_Territory", "British_Virgin_Islands",
		"Burkina_Faso", "Cape_Verde", "Cayman_Islands", "Central_African_Republic", "Christmas_Island", "Clipperton_Island",
		"Cocos_(Keeling)_Islands", "Congo,Democratic_Republic_of_the", "Congo,Republic_of_the", "Cook_Islands", "Coral_Sea_Islands",
		"Costa_Rica", "Cote_d'Ivoire", "Czech_Republic", "Dominican_Republic", "El_Salvador", "Equatorial_Guinea", "Europa_Island",
		"Falkland_Islands_(Islas_Malvinas)", "Faroe_Islands", "French_Guiana", "French_Polynesia", "French_Southern_and_Antarctic_Lands",
		"Gambia,The", "Gaza_Strip", "Glorioso_Islands", "Heard_Island_and_McDonald_Islands", "Holy_See_(Vatican_City)", "Hong_Kong",
		"Isle_of_Man", "Jan_Mayen", "Juan_de_Nova_Island", "Korea,North", "Korea,South", "Marshall_Islands", "Micronesia,Federated_States_of",
		"Navassa_Island", "Netherlands_Antilles", "New_Caledonia", "New_Zealand", "Norfolk_Island", "Northern_Mariana_Islands",
		"Papua_New_Guinea", "Paracel_Islands", "Pitcairn_Islands", "Puerto_Rico", "Saint_Helena", "Saint_Kitts_and_Nevis", "Saint_Lucia",
		"Saint_Pierre_and_Miquelon", "Saint_Vincent_and_the_Grenadines", "San_Marino", "Sao_Tome_and_Principe", "Saudi_Arabia",
		"Serbia_and_Montenegro", "Sierra_Leone", "Solomon_Islands", "South_Africa", "South_Georgia_and_the_South_Sandwich_Islands",
		"Spratly_Islands", "Sri_Lanka", "Trinidad_and_Tobago", "Tromelin_Island", "Turks_and_Caicos_Islands", "United_Arab_Emirates",
		"United_Kingdom", "United_States", "Virgin_Islands", "Wake_Island", "Wallis_and_Futuna", "West_Bank", "Western_Sahara",
	}
)

func main() {
	model := loadModel("../data/trained_model.txt", dim)
	f, err := os.Create("../data/q96.out.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	vCountries := make(map[string][]float64, 0)

	for _, c := range countries {
		if len(model[c]) != 0 {
			vCountries[c] = model[c]
		}
	}

	enc := gob.NewEncoder(f)
	err = enc.Encode(vCountries)
	if err != nil {
		log.Fatal(err)
	}
}

func loadModel(file string, dim int) map[string][]float64 {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReaderSize(f, 4096)

	vec := make(map[string][]float64, 0)

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		str := strings.Split(string(l), " ")
		if len(str) != dim+2 {
			continue
		}

		data := make([]float64, dim)
		for i := 1; i < dim+1; i++ {
			data[i-1], err = strconv.ParseFloat(str[i], 64)
			if err != nil {
				panic(err)
			}
		}
		vec[str[0]] = data
	}
	return vec
}
