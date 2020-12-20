package day9

import (
	"strconv"
	"bufio"
	"os"
	"fmt"
)

type instruction struct {
	command string
	amount int
	alreadyHit bool
}

func Solve() {
	file, error := os.Open("/Users/maja/gitRepo/adventOfCode/day9/input.txt")

	if error != nil {
		fmt.Println("Could not read file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int

	for key := 0; scanner.Scan(); key++ {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}
	
	firstOutOfOrder := findFirstIncorrectCypher(numbers)
	sumOfMinAndMax := sumOfSmallestAndBiggestNumberOfContiguousRange(numbers)

	fmt.Println("first incorrect number", firstOutOfOrder)
	fmt.Println("sum of min and max", sumOfMinAndMax)
}

func findFirstIncorrectCypher(numbers []int) int {
	for i := 25 ; i < len(numbers) ; i++ {
		sumFound := false
		for j := i-25 ; j < i ; j++ {
			for k := i-25 ; k < i ; k++ {
				if numbers[j]+numbers[k] == numbers[i] && numbers[j] != numbers[k] {
					sumFound = true
				}
			}
		}
		if !sumFound {
			return numbers[i]
		}
	}
	return 1
}

func sumOfSmallestAndBiggestNumberOfContiguousRange(numbers []int) int {
	invalidNumber := findFirstIncorrectCypher(numbers)

	for i := range numbers {
		sum := 0
		endOfRange := 0
		for j := i ; sum < invalidNumber ; j++ {
			sum += numbers[j]
			endOfRange = j
		}

		if sum == invalidNumber {
			min := findMin(numbers[i:endOfRange+1])
			max := findMax(numbers[i:endOfRange+1])
			return min+max
		}
	}

	return 1
}

func findMin(n []int) int {
	m := n[0]
	for i := 0 ; i < len(n) ; i++ {
		if n[i] < m {
			m = n[i]
		}
	}
	return m
}

func findMax(n []int) int {
	m := n[0]
	for i := 0 ; i < len(n) ; i++ {
		if n[i] > m {
			m = n[i]
		}
	}
	return m
}