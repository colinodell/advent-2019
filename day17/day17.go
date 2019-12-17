package main

import (
	"advent-2019/intcode"
	"advent-2019/utils"
	"fmt"
)

func main() {
	v := PowerOnRobot()

	fmt.Println("----- Part 1 -----")
	fmt.Printf("The sum of alignment parameters is: %d\n\n", v.CalculateSumOfAlignmentPatterns())
}

type pixel int
const (
	scaffold pixel = '#'
	space pixel = '.'
	robot_north pixel = '^'
	robot_east pixel = '>'
	robot_south pixel = 'v'
	robot_west pixel = '<'
)

type CameraMap map[utils.Vector2]pixel

type VacuumRobot struct {
	computer *intcode.Intcode
	cameraMap map[utils.Vector2]pixel
}

func PowerOnRobot() *VacuumRobot {
	v := new(VacuumRobot)
	v.cameraMap = make(CameraMap)

	program := utils.ReadCSVNumbers("./day17/input.txt")
	v.computer = intcode.CreateIntcodeComputer(program...)

	x, y := 0, 0
	for _, p := range v.computer.Run() {
		if p == 10 {
			x, y = 0, y+1
			continue
		}

		pos := utils.Vector2{X:x, Y:y}
		v.cameraMap[pos] = pixel(p)

		x++
	}

	return v
}

func (v *VacuumRobot) CalculateSumOfAlignmentPatterns() int {
	otherDirections := []utils.Vector2{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	sum := 0
	for pos, p := range v.cameraMap {
		if p == scaffold {
			// Check all 4 sides
			surroundedByScaffolds := true
			for _, dir := range otherDirections {
				if neighbor, ok := v.cameraMap[pos.Add(dir)]; !ok || neighbor != scaffold {
					surroundedByScaffolds = false
					break
				}
			}

			if surroundedByScaffolds {
				sum += pos.X * pos.Y
			}
		}
	}

	return sum
}
