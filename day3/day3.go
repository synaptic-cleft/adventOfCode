package day3

import (
	"bufio"
	"fmt"

	helper "github.com/synaptic-cleft/adventOfCode/internal"
)

func Solve() {
	file := helper.GetInput("day3")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	trees := calculateTrees(lines, 1, true) * calculateTrees(lines, 3, true) * calculateTrees(lines, 5, true) * calculateTrees(lines, 7, true) * calculateTrees(lines, 1, false)

	fmt.Println("multiplied trees", trees)
}

func calculateTrees(lines []string, stepX int, everyRow bool) int {
	var x int = 0
	var trees int
	for i, rowContents := range lines[1:] {
		if everyRow || i % 2 == 1 {
			x+=stepX
			if (rowContents[x % len(rowContents)] == '#') {
				trees++
			}
		}
	}

	fmt.Println("amount of trees", trees)
	return trees
}