package main

import (
	"advent-2019/intcode"
	"advent-2019/utils"
	"fmt"
)

func main() {
	program := utils.ReadCSVNumbers("./day15/input.txt")
	droid := NewRepairDroid(program)

	fmt.Println("----- Part 1 -----")
	distance := droid.TraverseMaze(true)
	fmt.Printf("The shortest path is %d steps\n\n", distance)

	fmt.Println("----- Part 2 -----")
	droid.ResetMaze()
	longestDistance := droid.TraverseMaze(true)
	fmt.Printf("The oxygen will take %d minutes to fill the area\n\n", longestDistance)
}

type direction int
type tile int

const (
	_ direction = iota
	North
	South
	West
	East
)

const(
	Unexplored tile = iota
	Wall
	Corridor
	OxygenSystem
)

var (
	Up = utils.Vector2{X:0, Y:-1}
	Down = utils.Vector2{X:0, Y:1}
	Left = utils.Vector2{X:-1, Y:0}
	Right = utils.Vector2{X:1, Y:0}
)

var (
	directions = map[direction]utils.Vector2{North:Up, South:Down, West:Left, East:Right}
	opposites = map[direction]direction{North:South, South:North, West:East, East:West}
)

type Path []direction

type RepairDroid struct {
	computer intcode.Intcode
	input, output chan int
	pos utils.Vector2
	maze map[utils.Vector2]tile
}

func NewRepairDroid(program []int) *RepairDroid {
	input := make(chan int)

	r := RepairDroid{
		input: input,
		output: intcode.CreateIntcodeComputer(program...).RunAsync(input, nil),
		maze: make(map[utils.Vector2]tile),
	}

	return &r
}

func (r *RepairDroid) Move(dir direction) bool {
	// Instruct computer to go there
	r.input <- int(dir)

	// Update our knowledge of the room (and our position too, if we moved)
	attemptedPos := r.pos.Add(directions[dir])
	switch <-r.output {
	case 0:
		r.maze[attemptedPos] = Wall
		return false
	case 1:
		r.pos = attemptedPos
		r.maze[r.pos] = Corridor
		return true
	case 2:
		r.pos = attemptedPos
		r.maze[r.pos] = OxygenSystem
		return true
	}

	panic("Invalid output")
}

func (r *RepairDroid) TraverseMaze(stopWhenOxygenFound bool) int {
	var pathsQueue []Path
	pathsQueue = append(pathsQueue, Path{})
	farthestDistance := 0

	for len(pathsQueue) > 0 {
		var path Path
		var history Path
		path, pathsQueue = pathsQueue[0], pathsQueue[1:]
		distance := 0

		// Move along this path
		navigationSuccessful := true
		for _, dir := range path {
			// Tell robot to move
			if r.Move(dir) {
				distance++

				// Is the oxygen here?
				if r.maze[r.pos] == OxygenSystem && stopWhenOxygenFound {
					return distance
				}

				history = append(history, dir)
				farthestDistance = utils.Max(farthestDistance, distance)
			} else {
				navigationSuccessful = false
				break
			}
		}

		if navigationSuccessful {
			// Enqueue all possible directions from this point that have not been visited yet
			for dir, vector := range directions {
				newPos := r.pos.Add(vector)
				if _, ok := r.maze[newPos]; ok {
					continue
				}

				newPath := make(Path, len(path))
				copy(newPath, path)
				newPath = append(newPath, dir)

				pathsQueue = append(pathsQueue, newPath)
			}
		}

		// Go back to the starting point
		for i := range history {
			k := len(history)-1-i
			dir := history[k]
			r.Move(opposites[dir])
		}
	}

	return farthestDistance
}

func (r *RepairDroid) ResetMaze() {
	// Reset our knowledge of the maze but not our current position
	r.maze = make(map[utils.Vector2]tile)
}
