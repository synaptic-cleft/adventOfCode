package day6

import (
	"strings"
	"bufio"
	"fmt"

	helper "github.com/synaptic-cleft/adventOfCode/internal"
)

type instruction struct {
	command string
	amount int
	alreadyHit bool
}

func Solve() {
	file := helper.GetInput("day6")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sumOfAnyYes := calculateAnyYes(lines)
	sumOfAllYes := calculateAllYes(lines)

	fmt.Println("sum of any yesses", sumOfAnyYes)
	fmt.Println("sum of all yesses", sumOfAllYes)
}

func calculateAnyYes(lines []string) int {
	totalCountOfYesses := 0
	stringsPerGroup := []string{}
	groupNo := 0
	for _, a := range lines {
		if a == "" {
			groupNo++
		} else if len(stringsPerGroup) == groupNo {
			stringsPerGroup = append(stringsPerGroup, a)
		} else {
			stringsPerGroup[groupNo] = stringsPerGroup[groupNo] + a
		}
	}
	
	allQuestions := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
						"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, group := range stringsPerGroup {
		for _, question := range allQuestions {
			if strings.Contains(group, question) {
				totalCountOfYesses++
			}
		}
	}

	return totalCountOfYesses
}

func calculateAllYes(lines []string) int {
	totalCountOfYesses := 0
	stringsPerGroup := [][]string{}
	groupNo := 0
	for _, a := range lines {
		if a == "" {
			groupNo++
		} else if len(stringsPerGroup) == groupNo {
			stringsPerGroup = append(stringsPerGroup, []string{a})
		} else {
			stringsPerGroup[groupNo] = append(stringsPerGroup[groupNo], a)
		}
	}
	
	allQuestions := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
						"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, group := range stringsPerGroup {
		for _, question := range allQuestions {
			if allStringsContain(question, group) {
				totalCountOfYesses++
			}
		}
	}

	return totalCountOfYesses
}

func allStringsContain(question string, group []string) bool {
	for _, v := range group {
		if !strings.Contains(v, question) {
			return false
		}
	}
	return true
}