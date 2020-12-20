package day5

import (
	"strconv"
	"sort"
	"math"
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
	file, error := os.Open("/Users/maja/gitRepo/adventOfCode/day5/input.txt")

	if error != nil {
		fmt.Println("Could not read file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var seats []string
	
	for scanner.Scan() {
		seats = append(seats, scanner.Text())
	}

	fmt.Println("highest boarding pass number", getHighestBoardingPassNumber(seats))

	// haha, in hindsight, i could have done this less complicated
	fmt.Println("the empty seat has ID", getBoardingPassNumberOfEmptySeat(seats))
}

func getHighestBoardingPassNumber(seats []string) int {
	allBoardingPassNumbers := []int{}

	for _, v := range seats {
		allBoardingPassNumbers = append(allBoardingPassNumbers, calculateBoardingPassNumber(v))
	}

	return findMax(allBoardingPassNumbers)
}

func calculateBoardingPassNumber(s string) int {
	boardingNumber := 0
	for k, v := range s {
		if v == 'B' {
			boardingNumber += int(math.Exp2(6-float64(k)))
		}
	}
	boardingNumber = boardingNumber * 8
	for k, v := range s[7:] {
		if v == 'R' {
			boardingNumber += int(math.Exp2(float64(2-k)))
		}
	}
	return boardingNumber
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

func getBoardingPassNumberOfEmptySeat(seats []string) int {
	allBoardingPassNumbers := []int{}

	for _, v := range seats {
		allBoardingPassNumbers = append(allBoardingPassNumbers, calculateLogicalBoardingPassNumber(v))
	}
	logicalNumber := findGap(allBoardingPassNumbers)
	originalSeatNumber := translateBackToSeatNumber(logicalNumber)
	actualBoardingPassNumber := calculateBoardingPassNumber(originalSeatNumber)

	return actualBoardingPassNumber
}

func calculateLogicalBoardingPassNumber(s string) int {
	boardingNumber := 0
	for k, v := range s {
		if v == 'B' {
			boardingNumber += int(math.Exp2(6-float64(k)))
		}
	}
	boardingNumber = boardingNumber *10
	for k, v := range s[7:] {
		if v == 'R' {
			boardingNumber += int(math.Exp2(float64(2-k)))
		}
	}
	return boardingNumber
}

func findGap(passes []int) int {
	sort.Ints(passes)
	for i, v := range passes {
		if passes[i+1] != v + 1 && passes[i+1] % 10 != 0 {
			return v+1
		}
	}
	fmt.Println(passes)
	return 0
}

func translateBackToSeatNumber(logicalNumber int) string {
	seatNumber := ""

	nr := logicalNumber /10
	for i := 6; i>=0 ; i-- {
		if (nr / int(math.Exp2(float64(i)))) >= 1 {
			seatNumber += "B"
			nr -= int(math.Exp2(float64(i)))
		} else {
			seatNumber += "F"
		}
	}

	lastDigit, _ := strconv.Atoi(strconv.Itoa(logicalNumber)[len(strconv.Itoa(logicalNumber))-1:])
	for i := 2; i>=0 ; i-- {
		if (lastDigit / int(math.Exp2(float64(i)))) >= 1 {
			seatNumber += "R"
			lastDigit -= int(math.Exp2(float64(i)))
		} else {
			seatNumber += "L"
		}
	}

	return seatNumber
}