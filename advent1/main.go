package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("InputProblem1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	input := make([]int, 0, len(records))
	for _, v := range records {
		n, err := strconv.Atoi(v[0])
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, n)
	}
	output, err := expenseReportFixer(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}

func expenseReportFixer(input []int) (int, error) {

	for i, v := range input {

		for _, w := range input[i+1:] {
			if v+w == 2020 {
				return v * w, nil
			}
		}
	}

	return 0, fmt.Errorf("no combination found")
}
