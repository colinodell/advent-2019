package main

import (
	"advent-2019/utils"
	"fmt"
)

func main() {
	program := utils.ReadCSVNumbers("./day02/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Result: %d\n", emulate(program, 12, 2))

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			if emulate(program, noun, verb) == 19690720 {
				fmt.Println("----- Part 2 -----")
				fmt.Printf("Result: %d\n", 100 * noun + verb)
				return
			}
		}
	}
}

func emulate (program []int, noun int, verb int) (result int) {
	memory := make([]int, len(program))
	copy(memory, program)

	memory[1], memory[2] = noun, verb

	for pos := 0 ; ; {
		opcode := memory[pos]
		switch opcode {
		case 1: // Addition
			op1Pos, op2Pos, resPos := memory[pos + 1], memory[pos + 2], memory[pos + 3]
			memory[resPos] = memory[op1Pos] + memory[op2Pos]

			pos += 4
		case 2: // Multiplication
			op1Pos, op2Pos, resPos := memory[pos + 1], memory[pos + 2], memory[pos + 3]
			memory[resPos] = memory[op1Pos] * memory[op2Pos]

			pos += 4
		case 99:
			return memory[0]
		}
	}
}
