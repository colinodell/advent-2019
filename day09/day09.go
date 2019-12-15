package main

import (
	"advent-2019/intcode"
	"advent-2019/utils"
	"fmt"
)

func main() {
	program := utils.ReadCSVNumbers("./day09/input.txt")

	fmt.Println("----- Part 1 -----")
	i := intcode.CreateIntcodeComputer(program...)
	outputs := i.Run(1)
	fmt.Printf("BOOST keycode (test mode): %d\n", outputs[0])

	fmt.Println("----- Part 2 -----")
	outputs = i.Run(2)
	fmt.Printf("BOOST keycode (boost mode): %d\n", outputs[0])
}
