package test

import (
	"fmt"
	"learninggo/Utils"
	"learninggo/adventofcode2"
)

func main() {
	fmt.Println("Fetching Correct Number of Passwords.....")
	var inputPasswordPolicy []string
	inputPasswordPolicy = Utils.IntializeInputString("InputProblem2.txt")
	fmt.Println("The total number of valid passwords by count is")
	fmt.Println(adventofcode2.CheckPasswordOccurenceCount(inputPasswordPolicy))
	fmt.Println("The total number of valid passwords by position is")
	fmt.Println(adventofcode2.CheckPasswordPosition(inputPasswordPolicy))

}
