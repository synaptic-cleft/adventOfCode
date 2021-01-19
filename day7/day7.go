package day7

import (
	"strconv"
	"regexp"
	"strings"
	"bufio"
	"fmt"

	helper "github.com/synaptic-cleft/adventOfCode/internal"
)

var bagColor int = 0

func Solve() {
	file := helper.GetInput("day7")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	m := make(map[string]string)

	for _, value := range lines {
		values := strings.Split(value, "contain")
		m[values[0]] = values[1]
	}

	allBagColors := findBags(m, "shiny gold", nil)

	// fmt.Println("different coloredBags", unique(allBagColors))
	fmt.Println("amount of different coloredBags", len(unique(allBagColors)))

	numberOfBagsInsideGoldenBag := calculateBagsInside(m, 1, "shiny gold") -1
	fmt.Println("amount of Bags inside shiny gold bag", numberOfBagsInsideGoldenBag)
}

func findBags(bagMap map[string]string, pattern string, alreadyUsedColors []string) []string {
	for key, value := range bagMap {
		if strings.Contains(value, pattern) {
			newPattern := key[:len(key)-2]
			alreadyUsedColors = findBags(bagMap, newPattern, append(alreadyUsedColors, newPattern))
		}
	}
	return alreadyUsedColors
}

func unique(array []string) []string {
    keys := make(map[string]bool)
    list := []string{} 
    for _, entry := range array {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

func calculateBagsInside(bagMap map[string]string, currentAmountOfBags int64, bagType string) int64 {
	bagsInside := getBagsMap(bagMap[bagType + " bags "])

	if len(bagsInside) == 0 {
		return currentAmountOfBags
	}

	innerBags := []int64{currentAmountOfBags}
	for typeOfInnerBag, no := range bagsInside {
		innerBags = append(innerBags, calculateBagsInside(bagMap, currentAmountOfBags * no, typeOfInnerBag)) 
	}

	var totalAmountOfBags int64
	for _, v := range innerBags {  
		totalAmountOfBags += v  
	   }

	return totalAmountOfBags
}

func getBagsMap(containingString string) map[string]int64 {
	m := make(map[string]int64)
	bags := strings.Split(containingString, ",")
	regex := regexp.MustCompile(`(\d)\s([a-z\s]*)\sbag`)

	for i := range bags {
		matches := regex.FindStringSubmatch(bags[i])
		if matches != nil {
			no, _ := strconv.ParseInt(matches[1], 10, 64)
			m[matches[2]] = no
		}
	}

	return m
}