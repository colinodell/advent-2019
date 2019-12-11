package main

import (
	"advent-2019/intcode"
	"advent-2019/utils"
	"fmt"
)

func main() {
	program := utils.ReadCSVNumbers("./day05/input.txt")

	fmt.Println("----- Part 1 -----")
	computer := intcode.CreateIntcodeComputer(program...)
	outputs := computer.Run(1)

	diagnosticCode := outputs[len(outputs)-1]

	fmt.Printf("Diagnostic code: %d\n", diagnosticCode)
}
