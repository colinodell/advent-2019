package main

import (
	"advent-2019/utils"
	"fmt"
)

func main() {
	modules := utils.ReadNumbers("./day01/input.txt")

	baseFuel := 0
	totalFuel := 0

	for _, mass := range modules {
		baseFuel += calculateFuelRequirement(mass)
		totalFuel += calculateTotalFuelRequirement(mass)
	}

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Base fuel: %d\n", baseFuel)

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Total fuel: %d\n", totalFuel)
}

func calculateFuelRequirement(mass int) int {
	return (mass / 3) - 2
}

func calculateTotalFuelRequirement(mass int) int {
	totalFuel := 0
	additionalWeight := mass

	for {
		additionalFuel := calculateFuelRequirement(additionalWeight)

		if additionalFuel <= 0 {
			break
		}

		totalFuel += additionalFuel
		additionalWeight = additionalFuel
	}

	return totalFuel
}