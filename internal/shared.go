package internal

import (
	"fmt"
	"os"
)

// GetInput Open input file
func GetInput(day string) *os.File {
	file, error := os.Open("/Users/maja/gitRepo/adventOfCode/" + day + "/input.txt")

	if error != nil {
		fmt.Println("Could not read file.")
		os.Exit(1)
	}

	return file
}