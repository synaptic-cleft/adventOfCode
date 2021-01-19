package day10

import (
	"sort"
	"strconv"
	"bufio"
	"fmt"

	helper "github.com/synaptic-cleft/adventOfCode/internal"
)

func Solve() {
	file := helper.GetInput("day10")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []int
	
	for scanner.Scan() {
		adapter, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, adapter)
	}

	result := calculateDifferences(lines)

	fmt.Println("difference in adapters multiplied", result)

	distinctNumberOfArrangements := calculateDistinctArrangements(lines)

	fmt.Println("difference in adapters multiplied", distinctNumberOfArrangements)

}

func calculateDifferences(adapters []int) int {
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	diffs := map[int]int{}
	for i, v := range adapters {
		if i != 0 {
			diff := v-adapters[i-1]
			diffs[diff]++
		} else {
			diffs[v]=v
		}
	}

	return diffs[1]*diffs[3]
}

func calculateDistinctArrangements(adapters []int) int {
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	
	return 0
}