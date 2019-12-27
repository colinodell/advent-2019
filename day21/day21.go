package main

import (
	"advent-2019/intcode"
	"advent-2019/utils"
	"fmt"
)

const program1 = `NOT C T
NOT A J
OR T J
AND D J
WALK
`

const program2 = `NOT E T
NOT H J
AND T J
NOT J J
NOT C T
AND T J
NOT B T
OR T J
NOT A T
OR T J
AND D J
RUN
`

func main() {
	brains := utils.ReadCSVNumbers("./day21/input.txt")
	i := intcode.CreateIntcodeComputer(brains...)

	fmt.Println("----- Part 1 -----")
	result, output := i.RunAscii(program1)
	fmt.Println(output)
	fmt.Println(result)

	fmt.Println("----- Part 2 -----")
	result, output = i.RunAscii(program2)
	fmt.Println(output)
	fmt.Println(result)
}
