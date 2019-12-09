package main

import (
	"advent-2019/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	wires := utils.ReadLines("./day03/input.txt")

	part1, part2 := Solve(wires[0], wires[1])

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Closest intersection to central port: %d\n", part1)

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Shortest distance to intersection: %d\n", part2)
}

func parseDirection (path string) (string, int) {
	return path[0:1], utils.ToInt(path[1:])
}

func moveOneUnit(pos Point, direction string) Point {
	switch direction {
	case "U":
		return Point{X: pos.X, Y: pos.Y - 1}
	case "D":
		return Point{X: pos.X, Y: pos.Y + 1}
	case "L":
		return Point{X: pos.X - 1, Y: pos.Y}
	case "R":
		return Point{X: pos.X + 1, Y: pos.Y}
	}

	panic("Invalid direction")
}

func Solve(wire1, wire2 string) (int, int) {
	grid := make(map[Point]int)

	// Travel the first wire
	pos, distance := Point{}, 0
	for _, path := range strings.Split(wire1, ",") {
		for direction, length := parseDirection(path); length > 0 ; length-- {
			pos = moveOneUnit(pos, direction)
			distance++
			grid[pos] = distance
		}
	}

	// Travel the second wire, but look for intersections as we go
	closestIntersection := math.MaxInt32
	shortestDistance := math.MaxInt32

	pos, distance = Point{}, 0
	for _, path := range strings.Split(wire2, ",") {
		for direction, length := parseDirection(path); length > 0 ; length-- {
			pos = moveOneUnit(pos, direction)
			distance++
			if grid[pos] > 0 {
				manhattanDistance := pos.ManhattanDistance()
				if manhattanDistance < closestIntersection {
					closestIntersection = manhattanDistance
				}
				combinedDistance := distance + grid[pos]
				if combinedDistance < shortestDistance {
					shortestDistance = combinedDistance
				}
			}
		}
	}

	return closestIntersection, shortestDistance
}

type Point struct {
	X int
	Y int
}

func (p Point) ManhattanDistance() int {
	return abs(p.X) + abs(p.Y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
