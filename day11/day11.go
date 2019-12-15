package main

import (
	"advent-2019/intcode"
	"advent-2019/utils"
	"fmt"
)

func main() {
	program := utils.ReadCSVNumbers("./day11/input.txt")

	fmt.Println("----- Part 1 -----")
	panels := runRobot(program, 0)
	fmt.Printf("Painted %d panels\n\n", len(panels))

	fmt.Println("----- Part 2 -----")
	panels = runRobot(program, 1)
	display(panels)

}

func runRobot(program []int, startingColor int) map[utils.Vector2]int {
	panels := make(map[utils.Vector2]int)

	up := utils.Vector2{X: 0, Y: -1}
	down := utils.Vector2{X: 0, Y: 1}
	left := utils.Vector2{X: -1, Y: 0}
	right := utils.Vector2{X: 1, Y: 0}

	pos := utils.Vector2{X: 0, Y: 0}
	direction := up

	panels[pos] = startingColor

	i := intcode.CreateIntcodeComputer(program...)
	done := make(chan struct{})
	input := make(chan int)
	output := i.RunAsync(input, func(){ close(done) })

	for {
		select {
		case input <- panels[pos]:
			color, turn := <-output, <-output

			panels[pos] = color

			if 0 == turn {
				// Turn left 90 degrees
				switch direction {
				case up:
					direction = left
				case left:
					direction = down
				case down:
					direction = right
				case right:
					direction = up
				}
			} else {
				// Turn right 90 degrees
				switch direction {
				case up:
					direction = right
				case right:
					direction = down
				case down:
					direction = left
				case left:
					direction = up
				}
			}

			pos = pos.Add(direction)
			case <-done:
				close(input)
				return panels
		}
	}
}

func display(panels map[utils.Vector2]int) {
	// Figure out the size we need to draw
	var min utils.Vector2
	var max utils.Vector2
	for v, _ := range panels {
		min = v.Min(min)
		max = v.Max(max)
	}

	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			pos := utils.Vector2{X:x, Y:y}
			if color, ok := panels[pos]; ok && color == 1 {
				fmt.Print("■")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
