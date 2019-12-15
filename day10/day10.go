package main

import (
	"advent-2019/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	input := utils.ReadFile("./day10/input.txt")
	asteroids := loadMap(input)

	fmt.Println("----- Part 1 -----")
	best, numberSeen := CalculateBestLocation(asteroids)
	fmt.Printf("The best asteroid is the one at (%s) which can see %d other asteroids\n\n", best.String(), numberSeen)

	fmt.Println("----- Part 2 -----")
	destroyed := Destroy(best, asteroids)
	fmt.Printf("The 200th asteroid destroyed was the one at (%s)\n", destroyed[199].String())
}

type Asteroid struct {
	X, Y int
}

type DetectedAsteroid struct {
	X, Y int
	angle float64
	distance float64
}

func (a Asteroid) String() string {
	return fmt.Sprintf("%d, %d", a.X, a.Y)
}

func (a DetectedAsteroid) String() string {
	return fmt.Sprintf("%d, %d", a.X, a.Y)
}

type DetectedAsteroids []DetectedAsteroid
func (a DetectedAsteroids) Len() int           { return len(a) }
func (a DetectedAsteroids) Less(i, j int) bool { return a[i].distance < a[j].distance }
func (a DetectedAsteroids) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func loadMap(input string) []Asteroid {
	var asteroids []Asteroid

	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if char == '#' {
				asteroids = append(asteroids, Asteroid{X: x, Y: y})
			}
		}
	}

	return asteroids
}

func CalculateBestLocation (asteroids []Asteroid) (Asteroid, int) {
	bestAsteroidCount := 0
	var bestAsteroid Asteroid

	for _, a := range asteroids {
		targets := ScanAsteroids(a, asteroids)
		if len(targets) > bestAsteroidCount {
			bestAsteroid, bestAsteroidCount = a, len(targets)
		}
	}

	return bestAsteroid, bestAsteroidCount
}

func ScanAsteroids(from Asteroid, others []Asteroid) map[float64]DetectedAsteroids {
	targets := make(map[float64]DetectedAsteroids)

	for _, target := range others {
		if target == from {
			continue
		}

		dX, dY := target.X - from.X, target.Y - from.Y
		angle := 2.0*math.Pi - (math.Atan2(float64(dX), float64(dY)) + math.Pi)

		targets[angle] = append(targets[angle], DetectedAsteroid{
			X:  target.X,
			Y: target.Y,
			angle: angle,
			distance: math.Sqrt(float64(dX*dX) + float64(dY*dY)),
		})
	}

	// At each angle, sort the asteroids by distance
	for _, asteroids := range targets {
		sort.Sort(asteroids)
	}

	return targets
}

func Destroy(from Asteroid, others []Asteroid) []DetectedAsteroid {
	var destroyed []DetectedAsteroid

	targets := ScanAsteroids(from, others)

	// Create a slice of angles from the list of targets
	angles := make([]float64, 0, len(targets))
	for k, _ := range targets {
		angles = append(angles, k)
	}
	sort.Float64s(angles)

	for len(destroyed) < len(others) - 1 {
		for _, angle := range angles {
			if len(targets[angle]) == 0 {
				continue
			}

			// Pop the closest asteroid off the list and destroy it
			target, remaining := targets[angle][0], targets[angle][1:]
			destroyed = append(destroyed, target)
			targets[angle] = remaining
		}
	}

	return destroyed
}
