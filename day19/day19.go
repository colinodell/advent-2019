package main

import (
	"advent-2019/intcode"
	"advent-2019/utils"
	"fmt"
)

func main() {
	input := utils.ReadCSVNumbers("./day19/input.txt")
	computer := intcode.CreateIntcodeComputer(input...)

	fmt.Println("----- Part 1 -----")
	part1 := Part1(computer)
	fmt.Printf("Points affected in the 50x50 area: %d\n\n", part1)

	fmt.Println("----- Part 2 -----")
	part2 := Part2(computer)
	fmt.Printf("Santa's ship fits in the square at: %d\n", part2)
}

func Part1(computer *intcode.Intcode) int {
	count := 0

	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			out := computer.Run(x, y)
			if out[0] == 1 {
				fmt.Print("#")
				count++
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}


	return count
}

func Part2(computer *intcode.Intcode) int {
	// There's some weirdness where the tractor beam is missing in a couple rows, so skip those
	startX, startY := 0, 5
	for {
		// Find where the tractor beam begins in this row
		for {
			if computer.Run(startX, startY)[0] == 0 {
				startX++
			} else {
				break
			}
		}

		x, y := startX, startY

		// Look ahead to see if the ship will fit here
		// We already know about the top-left corner so check the other 3
		for {
			// Stop if we don't have enough x range
			if computer.Run(x+99, y)[0] == 0 {
				startY++
				break
			}

			if computer.Run(x, y+99)[0] == 1 && computer.Run(x+99,y+99)[0] == 1 {
				return x * 10000 + y
			}

			x++
		}
	}
}
