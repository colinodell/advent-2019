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

	messageOffset, _ := strconv.Atoi(input[:7])

	fmt.Println("----- Part 2 -----")
	fmt.Printf("The eight-digit message in the real signal is: %s\n\n", ProcessSignal(input, 100, 10000, messageOffset, 8))
}

func ProcessSignal(signal string, phases, repeatSignal, messageOffset, messageLength int) string {
	signal = strings.Repeat(signal, repeatSignal)

	signalSize := len(signal)

	if messageOffset > (signalSize / 2) {
		return RunFastFFT(signal, phases, messageOffset)[:messageLength]
	} else {
		return RunFFT(signal, phases)[messageOffset:messageLength]
	}
}

func extractSignal(signal []int, from, to int) []int {
	result := make([]int, to-from)
	for i := range result {
		result[i] = signal[(from+i)%len(signal)]
	}
	return result
}

func RunFastFFT(signal string, phases int, messageOffset int) string {
	signalSlice := stringToIntSlice(signal)
	signalSlice = extractSignal(signalSlice, messageOffset, len(signalSlice))

	for p := 0; p < phases; p++ {
		sum := 0
		for i := len(signalSlice) - 1; i >= 0; i-- {
			sum += signalSlice[i]
			signalSlice[i] = utils.Abs(sum) % 10
		}
	}

	return intSliceToString(signalSlice)
}

func RunFFT(signal string, phases int) string {
	signalSlice := stringToIntSlice(signal)
	signalSize := len(signalSlice)
	resultDigits := make([]int, signalSize)

	for p := 0; p < phases; p++ {
		for i := 0; i < signalSize; i++ {
			sum := 0
			for j := i; j < signalSize; j++ {
				switch getMultiplier(i+1, j+1) {
				case 0:
					continue
				case 1:
					sum += signalSlice[j] % 10
				case -1:
					sum -= signalSlice[j] % 10
				}
			}

			resultDigits[i] = utils.Abs(sum % 10)
		}

		signalSlice = resultDigits
	}

	return intSliceToString(signalSlice)
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

func stringToIntSlice(string string) []int {
	length := len(string)
	slice := make([]int, length)

	for i := 0; i < length; i++ {
		slice[i] = int(string[i]) - 48 // Quick and dirty ASCII math
	}

	return slice
}

func intSliceToString(slice []int) string {
	result := ""

	for _, digit := range slice {
		result += strconv.Itoa(digit)
	}

	return result
}
