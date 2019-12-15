package main

import (
	"advent-2019/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadFile("./day12/input.txt")

	fmt.Println("----- Part 1 -----")
	moons := loadMoons(input)
	SimulateSystem(moons, 1000)
	fmt.Printf("Total energy after 1000 steps: %d\n\n", moons.CalcEnergy())

	fmt.Println("----- Part 2 -----")
	moons = loadMoons(input)
	steps := SimulateSystemUntilMatchesInitialState(moons)
	fmt.Printf("It takes %d steps to return to the initial state\n\n", steps)
}

func SimulateSystem(moons LunarSystem, steps int) {
	for i := 0; i < steps; i++ {
		applyGravity(moons)
		applyVelocity(moons)
	}
}

func SimulateSystemUntilMatchesInitialState(moons LunarSystem) int {
	xSteps, ySteps, zSteps := 0, 0, 0

	for steps := 0; xSteps == 0 || ySteps == 0 || zSteps == 0; {
		applyGravity(moons)
		applyVelocity(moons)

		steps++

		// Have all moons reached their original position and velocity in each dimension?
		if xSteps == 0 {
			allMoonsAt0 := true
			for _, m := range moons {
				if m.velocity.X != 0 {
					allMoonsAt0 = false
					break
				}
			}
			if allMoonsAt0 {
				xSteps = steps
			}
		}

		if ySteps == 0 {
			allMoonsAt0 := true
			for _, m := range moons {
				if m.velocity.Y != 0 {
					allMoonsAt0 = false
					break
				}
			}
			if allMoonsAt0 {
				ySteps = steps
			}
		}

		if zSteps == 0 {
			allMoonsAt0 := true
			for _, m := range moons {
				if m.velocity.Z != 0 {
					allMoonsAt0 = false
					break
				}
			}
			if allMoonsAt0 {
				zSteps = steps
			}
		}
	}

	return 2 * utils.LCM(xSteps, ySteps, zSteps)
}

type LunarSystem []Moon

type Moon struct {
	position, velocity utils.Vector3
}

func loadMoons(input string) LunarSystem {
	var system LunarSystem
	re := regexp.MustCompile(`-?\d+`)
	for _, line := range strings.Split(input, "\n") {
		pos := re.FindAllString(line, -1)

		if pos == nil {
			break
		}

		x, _:= strconv.Atoi(pos[0])
		y, _:= strconv.Atoi(pos[1])
		z, _:= strconv.Atoi(pos[2])

		moon := Moon{position: utils.Vector3{X:x, Y:y, Z:z}}
		system = append(system, moon)
	}

	return system
}

func applyGravity(moons LunarSystem) {
	for i, a := range moons {
		for j, b := range moons {
			if i == j {
				continue
			}

			delta := a.position.Compare(b.position)
			a.velocity = a.velocity.Add(delta)
		}

		moons[i] = a
	}
}

func applyVelocity(moons LunarSystem) {
	for i, moon := range moons {
		moons[i].position = moon.position.Add(moon.velocity)
	}
}

func (m *Moon) CalcEnergy() int {
	potential := utils.Abs(m.position.X) + utils.Abs(m.position.Y) + utils.Abs(m.position.Z)
	kinetic := utils.Abs(m.velocity.X) + utils.Abs(m.velocity.Y) + utils.Abs(m.velocity.Z)

	return potential * kinetic
}

func (m *Moon) String() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>", m.position.X, m.position.Y, m.position.Z, m.velocity.X, m.velocity.Y, m.velocity.Z)
}

func (moons LunarSystem) CalcEnergy() int {
	energy := 0
	for _, m := range moons {
		energy += m.CalcEnergy()
	}

	return energy
}
