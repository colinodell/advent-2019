package main

import (
	"advent-2019/intcode"
	"advent-2019/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	v := PowerOnRobot()

	fmt.Println("----- Part 1 -----")
	fmt.Printf("The sum of alignment parameters is: %d\n\n", v.CalculateSumOfAlignmentPatterns())

	fmt.Println("----- Part 2 -----")
	path := v.CalculateNavigationPath()
	fmt.Printf("Navigation path: %s\n", path)

	routine, a, b, c := reducePath(path)
	fmt.Printf("Reduced to %s with: %v | %v | %v\n", routine, a, b, c)

	fmt.Printf("Final result: %d", v.WakeUpAndMove(routine, a, b, c))
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

func (v *VacuumRobot) WakeUpAndMove(routine, a, b, c string) int {
	v.computer.OverwriteProgram(0, 2)

	input := fromAscii(routine, a, b, c, "n")

	output := v.computer.Run(input...)

	return output[len(output)-1]
}

func (v *VacuumRobot) CalculateNavigationPath() string {
	// What direction is the robot currently facing?
	var currentDirection pixel
	var currentLocation utils.Vector2
	for pos, p := range v.cameraMap {
		if p != scaffold && p != space {
			currentLocation = pos
			currentDirection = p
			break
		}
	}

	var navigationPath string
	for {
		// Can we turn left or right from here?
		var leftVector utils.Vector2
		var rightVector utils.Vector2
		var leftDirection pixel
		var rightDirection pixel

		switch currentDirection {
		case robot_north:
			leftDirection = robot_west
			leftVector = utils.Vector2{X:-1, Y:0}
			rightDirection = robot_east
			rightVector = utils.Vector2{X:1, Y:0}
		case robot_east:
			leftDirection = robot_north
			leftVector = utils.Vector2{X:0, Y:-1}
			rightDirection = robot_south
			rightVector = utils.Vector2{X:0, Y:1}
		case robot_south:
			leftDirection = robot_east
			leftVector = utils.Vector2{X:1, Y:0}
			rightDirection = robot_west
			rightVector = utils.Vector2{X:-1, Y:0}
		case robot_west:
			leftDirection = robot_south
			leftVector = utils.Vector2{X:0, Y:1}
			rightDirection = robot_north
			rightVector = utils.Vector2{X:0, Y:-1}
		}

		var stepVector utils.Vector2
		left, okLeft := v.cameraMap[currentLocation.Add(leftVector)]
		right, okRight := v.cameraMap[currentLocation.Add(rightVector)]
		if okLeft && left == scaffold {
			// Turning left
			stepVector = leftVector
			currentDirection = leftDirection
			navigationPath += "L,"
		} else if okRight && right == scaffold {
			// Turning right
			stepVector = rightVector
			currentDirection = rightDirection
			navigationPath += "R,"
		} else {
			// We've reached the end
			break
		}

		// Follow this direction until the end
		for i := 1; ; i++ {
			testLocation := currentLocation.Add(stepVector.Multiply(i))
			if p, ok := v.cameraMap[testLocation]; !ok || p != scaffold {
				currentLocation = currentLocation.Add(stepVector.Multiply(i-1))
				navigationPath += strconv.Itoa(i-1) + ","
				break
			}
		}
	}

	return strings.Trim(navigationPath, ",")
}

func reducePath(path string) (string, string, string, string) {
	// Brute-force a solution for A, B, and C
	for i := 3; i <= 20; i++ {
		for j := 3; j <= 20; j++ {
			for k := 3; k <= 20; k++ {
				remaining := path

				a := remaining[0:i]
				remaining = regexp.MustCompile(a + ",?").ReplaceAllString(remaining, "")

				b := remaining[0:j]
				remaining = regexp.MustCompile(b + ",?").ReplaceAllString(remaining, "")

				c := remaining[0:k]
				remaining = regexp.MustCompile(c + ",?").ReplaceAllString(remaining, "")

				if remaining == "" {
					routine := strings.ReplaceAll(path, a, "A")
					routine = strings.ReplaceAll(routine, b, "B")
					routine = strings.ReplaceAll(routine, c, "C")

					return routine, a, b, c
				}
			}
		}
	}

	panic("No solution found")
}

func fromAscii(input ...string) []int {
	var result []int

	for _, str := range input {
		for _, char := range str {
			result = append(result, int(char))
		}
		result = append(result, '\n')
	}

	return result
}
