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
	fmt.Printf("The first 8 digits after 100 phases of FFT are: %s\n\n", RunFFT(input, 100)[:8])
}

func RunFFT(signal string, phases int) string {
	signalSize := len(signal)

	for p := 0; p < phases; p++ {
		var resultDigits []int

		for i := 0; i < signalSize; i++ {
			pattern := generatePattern(signalSize, i + 1)

			sum := 0
			for j := 0; j < signalSize; j++ {
				// Quick and dirty ASCII math
				d := int(signal[j]) - 48
				sum += (d * pattern[j]) % 10
			}

			resultDigits = append(resultDigits, utils.Abs(sum % 10))
		}

		signal = intSliceToString(resultDigits)
	}

	return signal
}

func generatePattern(size, position int) []int {
	basePattern := []int{0, 1, 0, -1}
	var result []int

	loop:
		for {
			for _, p := range basePattern {
				for j := 0; j < position; j++ {
					result = append(result, p)
					if len(result) > size {
						break loop
					}
				}
			}
		}

	return result[1:]
}

func intSliceToString(slice []int) string {
	result := ""

	for _, digit := range slice {
		result += strconv.Itoa(digit)
	}

	return result
}
