package Utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func IntializeInputInt(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var inputList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nextItem, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error in conversion", err)
		}
		inputList = append(inputList, nextItem)
	}
	return inputList
}

func IntializeInputString(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var inputList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputList = append(inputList, scanner.Text())
	}
	return inputList

}
