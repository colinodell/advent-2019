package main

import (
	"advent-2019/intcode"
	"advent-2019/utils"
	"fmt"
	"github.com/dbyio/heappermutations"
	"sync"
)

func main() {
	program := utils.ReadCSVNumbers("./day07/input.txt")

	fmt.Println("----- Part 1 -----")
	maxOutputSignal := 0
	for _, phaseSettings := range heappermutations.Ints([]int{0,1,2,3,4}) {
		out := runAmplifierChain(program, phaseSettings...)
		if out > maxOutputSignal {
			maxOutputSignal = out
		}
	}

	fmt.Printf("Highest output signal: %d\n", maxOutputSignal)

	fmt.Println("----- Part 2 -----")
	maxOutputSignal = 0
	for _, phaseSettings := range heappermutations.Ints([]int{5,6,7,8,9}) {
		out := runAmplifierChainWithFeedbackLoop(program, phaseSettings...)
		if out > maxOutputSignal {
			maxOutputSignal = out
		}
	}

	fmt.Printf("Highest output signal: %d\n", maxOutputSignal)
}

func runAmplifierChain(program []int, phaseSettings ...int) int {
	computer := intcode.CreateIntcodeComputer(program...)

	output := 0
	for i := 0; i < len(phaseSettings); i++ {
		out := computer.Run(phaseSettings[i], output)
		output = out[0]
	}

	return output
}

func runAmplifierChainWithFeedbackLoop(program []int, phaseSettings ...int) int {
	computers := make([]intcode.Intcode, len(phaseSettings))
	initialInput := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(len(phaseSettings))

	input := initialInput
	for i := 0; i < len(phaseSettings); i++ {
		computers[i] = intcode.CreateIntcodeComputer(program...)
		if i == len(phaseSettings) - 1 {
			computers[i].RunAsync(input, wg.Done)
			computers[i].SetOutput(initialInput)
			input <- phaseSettings[i]
		} else {
			tmp := input
			input = computers[i].RunAsync(input, wg.Done)
			tmp <- phaseSettings[i]
		}
	}

	initialInput <- 0

	wg.Wait()
	out := <- initialInput

	return out
}
