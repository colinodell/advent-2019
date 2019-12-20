package main

import (
	"advent-2019/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := utils.ReadLines("./day18/input.txt")
	maze := NewMaze(input)

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Shortest path length: %d\n\n", maze.FindShortestPath())
}

type direction int
const (
	_ direction = iota
	North
	South
	West
	East
)

var (
	Up = utils.Vector2{X:0, Y:-1}
	Down = utils.Vector2{X:0, Y:1}
	Left = utils.Vector2{X:-1, Y:0}
	Right = utils.Vector2{X:1, Y:0}

	directions = map[direction]utils.Vector2{North:Up, South:Down, West:Left, East:Right}
)

type Path struct {
	pos utils.Vector2
	steps int
	route []rune
}

func NewPath() Path {
	p := Path{}
	return p
}

func (p Path) Copy() Path {
	newPath := NewPath()

	newPath.pos = p.pos.Copy()
	newPath.steps = p.steps
	newPath.route = make([]rune, len(p.route))
	copy(newPath.route, p.route)

	return newPath
}

type Maze struct {
	raw []string
	tiles map[utils.Vector2]rune
	keys map[rune]utils.Vector2
}

func NewMaze(input []string) Maze {
	m := Maze{}
	m.raw = input
	m.tiles = make(map[utils.Vector2]rune)
	m.keys = make(map[rune]utils.Vector2)

	// Load the maze
	for y, line := range input {
		for x, char := range line {
			m.tiles[utils.Vector2{X: x, Y: y}] = char

			if char == '@' || (char >= 'a' && char <= 'z') {
				m.keys[char] = utils.Vector2{X: x, Y: y}
			}
		}
	}

	return m
}

func (m *Maze) FindShortestPath() int {
	paths := m.findPathsBetweenAllKeys()
	return m.reducePaths(paths)
}

func (m *Maze) findPathsBetweenAllKeys() map[rune]map[rune]Path {
	result := make(map[rune]map[rune]Path)

	for key1Name, key1Pos := range m.keys {
		result[key1Name] = make(map[rune]Path)

		for key2Name, key2Pos := range m.keys {
			if key1Name == key2Name || key2Name == '@' {
				continue
			}

			if _, ok := result[key2Name][key1Name]; ok {
				result[key1Name][key2Name] = result[key2Name][key1Name]
			}

			path := m.bfs(key1Pos, key2Pos)
			result[key1Name][key2Name] = path
		}
	}

	return result
}

func unlocked(doors []rune, keys []rune) bool {
	for _, door := range doors {
		if !runeSliceContains(door, keys) {
			return false
		}
	}

	return true
}

func runeSliceContains(needle rune, haystack []rune) bool {
	for _, x := range haystack {
		if strings.ToUpper(string(needle)) == strings.ToUpper(string(x)) {
			return true
		}
	}

	return false
}

type candidatePath struct {
	currentLocation rune
	keysObtained    int
}

func newCandidatePath(location rune, keysObtained []rune) candidatePath {
	ik := candidatePath{currentLocation: location, keysObtained: 0}

	for _, k := range keysObtained {
		ik.keysObtained |= 1 << (int(k) - 97)
	}


	return ik
}

func (ik *candidatePath) keysObtainedAsRuneSlice() []rune  {
	ret := make([]rune, 0, 26)

	for i := 0; i < 26; i++ {
		if (1 << i) & ik.keysObtained == (1 << i) {
			ret = append(ret, rune(i + 97))
		}
	}

	return ret
}

func (m *Maze) reducePaths(knownRoutes map[rune]map[rune]Path) int {
	var keys []rune
	for k, _ := range knownRoutes {
		if k >= 'a' && k <= 'z' {
			keys = append(keys, k)
		}
	}

	var candidatePaths = make(map[candidatePath]int, 0)
	candidatePaths[candidatePath{currentLocation: '@'}] = 0

	for range keys {
		nextCandidatePaths := make(map[candidatePath]int)
		for data, currentDistance := range candidatePaths {
			currentLocation, currentKeys := data.currentLocation, data.keysObtainedAsRuneSlice()
			for _, newKey := range keys {
				if !runeSliceContains(newKey, currentKeys) {
					routeInfo := knownRoutes[currentLocation][newKey]
					routeDistance, doors := routeInfo.steps, routeInfo.route
					reachable := unlocked(doors, currentKeys)
					if reachable {
						newCandidate := newCandidatePath(newKey, append(currentKeys, newKey))
						newDistance := currentDistance + routeDistance

						if _, ok := nextCandidatePaths[newCandidate]; !ok || (newDistance < nextCandidatePaths[newCandidate]) {
							nextCandidatePaths[newCandidate] = newDistance
						}
					}
				}
			}
		}

		candidatePaths = nextCandidatePaths
	}

	shortestDistance := math.MaxInt32
	for _, dist := range candidatePaths {
		if dist < shortestDistance {
			shortestDistance = dist
		}
	}

	return shortestDistance
}

func (m *Maze) bfs(from, to utils.Vector2) Path {
	var pathsQueue []Path
	pathsQueue = append(pathsQueue, Path{pos: from})

	alreadyVisited := make(map[utils.Vector2]bool)

	for i := 0; i < len(pathsQueue); i++ {
		path := pathsQueue[i]

		// At the given position, move forward in one direction until you hit something
		for _, vector := range directions {
			nextPos := path.pos.Add(vector)

			// We haven't been to that position, have we?
			if _, ok := alreadyVisited[nextPos]; ok {
				continue
			}

			alreadyVisited[nextPos] = true

			// Check for wall/boundry
			r, ok := m.tiles[nextPos]
			if !ok || r == '#' {
				continue
			}

			candidatePath := path.Copy()
			candidatePath.pos = nextPos
			candidatePath.steps++

			// Is there a door here?
			upperCase := strings.ToUpper(string(r))
			if r >= 'A' && r <= 'Z' && string(r) == upperCase {
				candidatePath.route = append(candidatePath.route, r)
			}

			// Are we at the location we were looking for?
			if candidatePath.pos == to {
				return candidatePath
			}

			pathsQueue = append(pathsQueue, candidatePath)
		}
	}

	panic("Could not find a path")
}
