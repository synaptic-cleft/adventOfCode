package day4

import (
	"regexp"
	"strconv"
	"strings"
	"bufio"
	"fmt"

	helper "github.com/synaptic-cleft/adventOfCode/internal"
)

func Solve() {
	file := helper.GetInput("day4")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	a := []string{}
	passportCounter := 0

	for _, value := range lines {
		if len(value) == 0 {
			passportCounter++
		} else {
			if len(a) > passportCounter {
				// fmt.Println(a)
				a[passportCounter] = a[passportCounter] + " " + value
			} else {
				// fmt.Println(a)
				a = append(a, value)
			}
		}
	}

	count := 0
	for _, value := range a {
		pairs := strings.Split(value, " ")

		m := make(map[string]string)

		for _, v := range pairs {
			onePair := strings.TrimSpace(v)
			leftAndRight := strings.Split(onePair, ":")
			m[leftAndRight[0]] = leftAndRight[1]
		}

		// if passportIsValid(m) {count++}
		if passportIsStrictlyValid(m) {count++}
		
	}

	fmt.Println(count)
}

func passportIsValid(attributes map[string]string) bool {
	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, v := range keys {
		if attributes[v] == "" {
			return false
		}
	}

	return true
}

func passportIsStrictlyValid(attributes map[string]string) bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	byr, e := strconv.Atoi(attributes["byr"])
	if e != nil {
		return false
	}

	if 2002 < byr || byr < 1920 {
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	iyr, e := strconv.Atoi(attributes["iyr"])
	if e != nil {
		return false
	}

	if 2020 < iyr || iyr < 2010 {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyr, e := strconv.Atoi(attributes["eyr"])
	if e != nil {
		return false
	}

	if 2030 < eyr || eyr < 2020 {
		return false
	}

	// hgt (Height) - a number followed by either cm or in:
	hgt := attributes["hgt"]
	matches, _ := regexp.MatchString(`^\d+(cm|in)$`, hgt)
	if !matches {
		return false
	}
	sizeNumber, _ := strconv.Atoi(hgt[:len(hgt)-2])
	// If cm, the number must be at least 150 and at most 193.
	if strings.Contains(hgt, "cm") && (sizeNumber < 150 || sizeNumber > 193) {
		return false
	}
	// If in, the number must be at least 59 and at most 76.
	if strings.Contains(hgt, "in") && (sizeNumber < 59 || sizeNumber > 76) {
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	colorMatch, _ := regexp.MatchString(`^#[\da-f]{6}$`, attributes["hcl"])
	if !colorMatch {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	switch attributes["ecl"] {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
	default:
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	pid, _ := regexp.MatchString(`^\d{9}$`, attributes["pid"])
	if !pid {
		return false
	}

	fmt.Println(attributes)
	return true
}
