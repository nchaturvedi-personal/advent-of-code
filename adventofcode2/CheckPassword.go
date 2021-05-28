package adventofcode2

import (
	"fmt"
	"strconv"
	"strings"
)

func CheckPasswordOccurenceCount(input []string) int {
	var validPasswordsCount int
	inputArraySize := len(input)
	for i := 0; i < inputArraySize; i++ {

		condition := input[i]
		minOccurs, err := strconv.Atoi(string(condition[0:strings.IndexByte(condition, '-')]))
		if err != nil {
			//fmt.Println(condition)
			fmt.Println("Error in conversion", err)
		}
		maxOccurs, err := strconv.Atoi(string(condition[strings.IndexByte(condition, '-')+1 : strings.IndexByte(condition, ' ')]))
		if err != nil {
			fmt.Println(condition)
			fmt.Println("Error in conversion", err)
		}
		character := string(condition[strings.IndexByte(condition, ':')-1])
		password := string(condition[strings.IndexByte(condition, ':')+2:])
		//fmt.Println(character)
		//fmt.Println(password)
		//fmt.Println(minOccurs)
		//fmt.Println(maxOccurs)
		actualOccurence := strings.Count(password, character)
		//fmt.Println(actualOccurence)

		if actualOccurence >= minOccurs && actualOccurence <= maxOccurs {
			validPasswordsCount = validPasswordsCount + 1
		}
	}

	return validPasswordsCount
}

func CheckPasswordPosition(input []string) int {

	var validPasswordsCount int
	inputArraySize := len(input)
	for i := 0; i < inputArraySize; i++ {

		condition := input[i]
		minOccurs, err := strconv.Atoi(string(condition[0:strings.IndexByte(condition, '-')]))
		if err != nil {
			//fmt.Println(condition)
			fmt.Println("Error in conversion", err)
		}
		maxOccurs, err := strconv.Atoi(string(condition[strings.IndexByte(condition, '-')+1 : strings.IndexByte(condition, ' ')]))
		if err != nil {
			fmt.Println(condition)
			fmt.Println("Error in conversion", err)
		}
		character := string(condition[strings.IndexByte(condition, ':')-1])
		password := string(condition[strings.IndexByte(condition, ':')+2:])
		//fmt.Println(character)
		//fmt.Println(password)
		//fmt.Println(minOccurs)
		//fmt.Println(maxOccurs)
		//actualOccurence := strings.Count(password, character)
		//fmt.Println(actualOccurence)

		if (string(password[minOccurs-1]) == character && string(password[maxOccurs-1]) != character) || string(password[minOccurs-1]) != character && string(password[maxOccurs-1]) == character {
			validPasswordsCount = validPasswordsCount + 1
		}
	}

	return validPasswordsCount
}
