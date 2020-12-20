package day1

import (
	"strconv"
	"bufio"
	"os"
	"fmt"
)

func Solve() {
	fmt.Println("hoi")

	// read input file
	// FIX: relative path
	file, error := os.Open("/Users/maja/gitRepo/adventOfCode/day1/input.txt")

	if error != nil {
		fmt.Println("Could not read file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []int

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, number)
	}

	// find two entries that some up to 2020
	twoNumbers := findTwoNumbersThatSumUpTo2020(lines)

	// multiply those two numbers
	fmt.Println(twoNumbers[0], twoNumbers[1], twoNumbers[0] * twoNumbers[1])

	// find three entries that some up to 2020
	threeNumbers := findThreeNumbersThatSumUpTo2020(lines)

	// multiply those three numbers
	fmt.Println(threeNumbers[0], threeNumbers[1], threeNumbers[2], threeNumbers[0] * threeNumbers[1] * threeNumbers[2])


}

func findTwoNumbersThatSumUpTo2020(lines []int) ([]int) {
	for _, number := range lines[1:] {
		if lines[0] + number == 2020 {
			return []int{lines[0], number}
		}
	}

	return findTwoNumbersThatSumUpTo2020(lines[1:])
}

func findThreeNumbersThatSumUpTo2020(lines []int) ([]int) {
	for _, n2 := range lines[1:] {
		for _, n3 := range lines[2:] {
			if lines[0] + n2 + n3 == 2020 {
				return []int{lines[0], n2, n3}
			}
		}
	}

	return findThreeNumbersThatSumUpTo2020(lines[1:])
}