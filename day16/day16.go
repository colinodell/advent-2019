package main

import (
	"advent-2019/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadFile("./day16/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("The first 8 digits after 100 phases of FFT are: %s\n\n", ProcessSignal(input, 100, 1, 0, 8))
}

func ProcessSignal(signal string, phases, repeatSignal, messageOffset, messageLength int) string {
	signal = strings.Repeat(signal, repeatSignal)

	processed := RunFFT(signal, phases)

	return processed[messageOffset:messageLength]
}

func RunFFT(signal string, phases int) string {
	signalSize := len(signal)
	resultDigits := make([]int, signalSize)

	for p := 0; p < phases; p++ {
		for i := 0; i < signalSize; i++ {
			sum := 0
			for j := i; j < signalSize; j++ {
				switch getMultiplier(i+1, j+1) {
				case 0:
					continue
				case 1:
					// Quick and dirty ASCII math
					sum += (int(signal[j]) - 48) % 10
				case -1:
					sum -= (int(signal[j]) - 48) % 10
				}
			}

			resultDigits[i] = utils.Abs(sum % 10)
		}

		signal = intSliceToString(resultDigits)
	}

	return signal
}

func getMultiplier(i, j int) int {
	switch (j / i) % 4 {
	case 0, 2:
		return 0
	case 1:
		return 1
	default:
		return -1
	}
}

func intSliceToString(slice []int) string {
	result := ""

	for _, digit := range slice {
		result += strconv.Itoa(digit)
	}

	return result
}
