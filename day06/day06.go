package main

import (
	"advent-2019/utils"
	"fmt"
	"strings"
)

func main() {
	orbits := loadOrbits(utils.ReadLines("./day06/input.txt"))

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Total number of direct and indirect orbits: %d\n\n", CountDirectAndIndirectOrbits(orbits))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Minimum number of orbital transfers: %d\n\n", CalculateOrbitalTransfers(orbits))
}

func CountDirectAndIndirectOrbits(orbits map[string]string) int {
	total := 0

	// For each object in our map...
	for object := range orbits {
		// Work backwards through its parents and count them...
		for {
			parentObject := orbits[object]

			object = parentObject
			total++

			// Stopping once we reach the end
			if object == "COM" {
				break
			}
		}
	}

	return total
}

func CalculateOrbitalTransfers(orbits map[string]string) int {
	youToCom := make(map[string]int)

	{
		// Starting at "YOU", count your way back up to COM
		object := orbits["YOU"]
		distance := 0
		for {
			parentObject := orbits[object]
			object = parentObject
			distance++
			youToCom[parentObject] = distance

			if object == "COM" {
				break
			}
		}
	}

	{
		// Starting at "SAN", count your way back up to COM,
		// at least until you reach an object already in youToCom
		object := orbits["SAN"]
		distance := 0
		for {
			object = orbits[object]
			distance++

			if _, ok := youToCom[object]; ok {
				return distance + youToCom[object]
			}
		}
	}
}

func loadOrbits(inputLines []string) map[string]string {
	orbitParentsByChild := make(map[string]string)

	for _, input := range inputLines {
		parent, child := splitLine(input)
		orbitParentsByChild[child] = parent
	}

	return orbitParentsByChild
}

func splitLine(line string) (string, string) {
	x := strings.SplitN(line, ")", 2)

	return x[0], x[1]
}
