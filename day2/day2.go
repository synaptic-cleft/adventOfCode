package day2

import (
	"strconv"
	"regexp"
	"strings"
	"bufio"
	"os"
	"fmt"
)


var correctPasswordCount = 0
var correctPasswordCountSecondRuleSet = 0

func Solve() {
	file, error := os.Open("./day2/input.txt")

	if error != nil {
		fmt.Println("Could not read file.")
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		amount, letterField, password := line[0], line[1], line[2]
		amounts := strings.Split(amount, "-")
		firstNumber, e1 := strconv.Atoi(amounts[0])
		secondNumber, e2 := strconv.Atoi(amounts[1])

		if e1 != nil || e2 != nil {
			fmt.Println("You promised the input would contain integers only. Why would you lie to me?")
			os.Exit(1)
		}

		letter := strings.Split(letterField, ":")[0]

		countCorrectPasswords(password, letter, firstNumber, secondNumber)
		countExactlyOnceOccurrence(password, []byte(letter)[0], firstNumber, secondNumber)
	}

	fmt.Println("correctPasswordCount", correctPasswordCount)
	fmt.Println("correctPasswordCountSecondRuleSet", correctPasswordCountSecondRuleSet)

}

func countCorrectPasswords(pw string, letter string, min int, max int) {
	regex := regexp.MustCompile(letter)

	occurenceOfNumberInPassword := len(regex.FindAllStringIndex(pw, -1))

	if (occurenceOfNumberInPassword >= min && occurenceOfNumberInPassword <= max) {
		correctPasswordCount++
	}
}

func countExactlyOnceOccurrence(pw string, letter byte, pos1 int, pos2 int) {
	if (pw[pos1-1] == letter && pw[pos2-1] != letter) || (pw[pos1-1] != letter && pw[pos2-1] == letter) {
		correctPasswordCountSecondRuleSet++
	}
}