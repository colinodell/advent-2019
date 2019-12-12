package main

import (
	"advent-2019/intcode"
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
	computer := intcode.CreateIntcodeComputer(program...)
	computer.ChangeNounAndVerb(noun, verb)
	computer.Run()

	return computer.Read(0)
}
