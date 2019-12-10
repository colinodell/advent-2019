package main

import (
	"fmt"
)

func main() {
	fmt.Println("----- Part 1 -----")
	fmt.Printf("Number of passwords meeting the criteria: %d\n", GenerateAndCountPasswords(134792, 675810, false))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Number of passwords meeting the criteria: %d\n", GenerateAndCountPasswords(134792, 675810, true))
}

func GenerateAndCountPasswords(min, max int, requireExactlyTwoSameDigits bool) int {
	count := 0
	for i := min; i <= max ; i++ {
		if PasswordMeetsCriteria(i, requireExactlyTwoSameDigits) {
			count++
		}
	}

	return count
}

func PasswordMeetsCriteria(password int, requireExactlyTwoSameDigits bool) bool {
	digits := getDigits(password)

	digitRuns := make(map[int]int, len(digits))
	runIndex := -1

	previousDigit := 0

	for _, d := range digits {
		if d < previousDigit {
			return false
		} else if d == previousDigit {
			digitRuns[runIndex]++
		} else {
			digitRuns[runIndex + 1] = 1
			runIndex++
		}

		previousDigit = d
	}

	for _, count := range digitRuns {
		if requireExactlyTwoSameDigits && count == 2 {
			return true
		} else if !requireExactlyTwoSameDigits && count >= 2 {
			return true
		}
	}

	return false
}

func getDigits(number int) []int {
	// Hard-coded for 6 digits
	ret := make([]int, 6)

	for i := len(ret) - 1 ; i >= 0 ; i-- {
		ret[i] = number % 10
		number /= 10
	}

	return ret
}
