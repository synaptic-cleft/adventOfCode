package day8

import (
	"strconv"
	"strings"
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
	file, error := os.Open("/Users/maja/gitRepo/adventOfCode/day8/input.txt")

	if error != nil {
		fmt.Println("Could not read file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructions := make(map[int]*instruction)

	for key := 0; scanner.Scan(); key++ {
		arr := strings.Split(scanner.Text(), " ")
		c := arr[0]
		a, _ := strconv.Atoi(arr[1])

		i := instruction{c, a, false}
		instructions[key] = &i
	}

	accumulator, error := playGame(instructions)

	if error != nil {
		fmt.Println(error.Error())
	}

	fmt.Println("accumulator", accumulator)

	terminateGame(instructions)
}

func playGame(instr map[int]*instruction) (int, error) {
	acc := 0
	for i := 0 ; i < len(instr) ;  {
		if instr[i].alreadyHit {
			return acc, nil
		}
		instr[i].alreadyHit = true

		if instr[i].command == "acc" {
			acc += instr[i].amount
			i++
		} else if instr[i].command == "jmp" {
			i += instr[i].amount
		} else {
			i++
		}
	}
	return acc, fmt.Errorf("No infinite loop")
}

func terminateGame(instr map[int]*instruction){
	for i := 0 ; i < len(instr) ; i++ {
		// copy map
		modifiedCopyOfInstr := map[int]*instruction{}
		for k,v := range instr {
			modifiedCopyOfInstr[k] = &instruction{v.command, v.amount, false}
		  }

		// change if jmp or nop
		if (instr[i].command == "jmp") {
			modifiedCopyOfInstr[i].command = "nop"
		} else if (instr[i].command == "nop") {
			modifiedCopyOfInstr[i].command = "jmp"
		}

		// play game with modified input
		acc, error := playGame(modifiedCopyOfInstr)
	
		if error != nil {
			fmt.Printf("Infinite loop repaired by changing index %d from %s to %s\n", 
			i, instr[i].command, modifiedCopyOfInstr[i].command)
			fmt.Println("Accumulator is", acc)
		}
	}
}