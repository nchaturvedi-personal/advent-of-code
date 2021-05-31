package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("InputProblem2.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	input := make([]string, 0, len(records))
	for _, v := range records {

		input = append(input, string(v[0]))

	}
	output, err := checkPasswords(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)

	outputWithNewPolicy, err := checkPasswordsNewPolicy(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(outputWithNewPolicy)

}

func checkPasswords(input []string) (int, error) {

	var validPasswords int

	for _, v := range input {
		minOccurs, err := strconv.Atoi(v[0:strings.IndexByte(v, '-')])
		if err != nil {
			log.Fatal()
		}
		_ = minOccurs
		maxOccurs, err := strconv.Atoi(v[strings.IndexByte(v, '-')+1 : strings.IndexByte(v, ' ')])
		if err != nil {
			log.Fatal()
		}
		character := string(v[strings.IndexByte(v, ' ')+1 : strings.IndexByte(v, ':')])
		password := string(v[strings.IndexByte(v, ':')+2:])
		if strings.Count(password, character) >= minOccurs && strings.Count(password, character) <= maxOccurs {
			validPasswords = validPasswords + 1
		}
	}

	return validPasswords, nil
}

func checkPasswordsNewPolicy(input []string) (int, error) {

	var validPasswords int

	for _, v := range input {
		firstPosition, err := strconv.Atoi(v[0:strings.IndexByte(v, '-')])
		if err != nil {
			log.Fatal()
		}
		_ = firstPosition
		secondPosition, err := strconv.Atoi(v[strings.IndexByte(v, '-')+1 : strings.IndexByte(v, ' ')])
		if err != nil {
			log.Fatal()
		}
		character := string(v[strings.IndexByte(v, ' ')+1 : strings.IndexByte(v, ':')])
		password := string(v[strings.IndexByte(v, ':')+2:])
		if (string(password[firstPosition-1]) == character && string(password[secondPosition-1]) != character) || (string(password[firstPosition-1]) != character && string(password[secondPosition-1]) == character) {
			validPasswords = validPasswords + 1

		}
	}

	return validPasswords, nil
}
