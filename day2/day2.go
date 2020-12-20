package day2

import (
	"strconv"
	"regexp"
	"strings"
	"bufio"
	"os"
	"fmt"
)


var correctPasswordCount int = 0
var correctPasswordCountSecondRuleSet int = 0

func Solve() {
	fmt.Println("hoi")

	// read input file
	// FIX: relative path
	file, error := os.Open("/Users/maja/gitRepo/adventOfCode/day2/input.txt")

	if error != nil {
		fmt.Println("Could not read file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		// lines = append(lines, scanner.Text())
		line := strings.Fields(scanner.Text())
		amount, letterField, password := line[0], line[1], line[2]
		amounts := strings.Split(amount, "-")
		firstNumber, _ := strconv.Atoi(amounts[0])
		secondNumber, _ := strconv.Atoi(amounts[1])
		letter := strings.Split(letterField, ":")[0]

		verifyIfMinAndMaxAmountOfConstraintIsMet(password, letter, firstNumber, secondNumber)
		verifyExactlyOnceOccurence(password, []byte(letter)[0], firstNumber, secondNumber)
	}

	fmt.Println("correctPasswordCount", correctPasswordCount)
	fmt.Println("correctPasswordCountSecondRuleSet", correctPasswordCountSecondRuleSet)

}

func verifyIfMinAndMaxAmountOfConstraintIsMet(pw string, letter string, min int, max int) {
	regex := regexp.MustCompile(letter)

	occurenceOfNumberInPassword := len(regex.FindAllStringIndex(pw, -1))

	if (occurenceOfNumberInPassword >= min && occurenceOfNumberInPassword <= max) {
		correctPasswordCount++
	}
}

func verifyExactlyOnceOccurence(pw string, letter byte, pos1 int, pos2 int) {
	if (pw[pos1-1] == letter && pw[pos2-1] != letter) || (pw[pos1-1] != letter && pw[pos2-1] == letter) {
		correctPasswordCountSecondRuleSet++
	}
}