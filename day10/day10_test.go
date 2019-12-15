package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadMap(t *testing.T) {
	input := `.#..#
.....
#####
....#
...##`
	asteroids := loadMap(input)

	assert.Len(t, asteroids, 10)

	assert.Contains(t, asteroids, Asteroid{X:1, Y:0})
	assert.Contains(t, asteroids, Asteroid{X:4, Y:0})
	assert.Contains(t, asteroids, Asteroid{X:0, Y:2})
	assert.Contains(t, asteroids, Asteroid{X:1, Y:2})
	assert.Contains(t, asteroids, Asteroid{X:2, Y:2})
	assert.Contains(t, asteroids, Asteroid{X:3, Y:2})
	assert.Contains(t, asteroids, Asteroid{X:4, Y:2})
	assert.Contains(t, asteroids, Asteroid{X:4, Y:3})
	assert.Contains(t, asteroids, Asteroid{X:3, Y:4})
	assert.Contains(t, asteroids, Asteroid{X:4, Y:4})
}

func TestCalculateBestLocation(t *testing.T) {
	input := `.#..#
.....
#####
....#
...##`
	asteroids := loadMap(input)
	bestAsteroid, numberSeen := CalculateBestLocation(asteroids)
	assert.Equal(t, Asteroid{X:3, Y:4}, bestAsteroid)
	assert.Equal(t, 8, numberSeen)

	input = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`
	asteroids = loadMap(input)
	bestAsteroid, numberSeen = CalculateBestLocation(asteroids)
	assert.Equal(t, Asteroid{X:11, Y:13}, bestAsteroid)
	assert.Equal(t, 210, numberSeen)
}

func TestDestroy(t *testing.T) {
	input := `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`
	asteroids := loadMap(input)
	bestAsteroid, _ := CalculateBestLocation(asteroids)
	destroyed := Destroy(bestAsteroid, asteroids)

	assert.Equal(t, "11, 12", destroyed[0].String())
	assert.Equal(t, "12, 1", destroyed[1].String())
	assert.Equal(t, "12, 2", destroyed[2].String())
	assert.Equal(t, "12, 8", destroyed[9].String())
	assert.Equal(t, "16, 0", destroyed[19].String())
	assert.Equal(t, "16, 9", destroyed[49].String())
	assert.Equal(t, "10, 16", destroyed[99].String())
	assert.Equal(t, "9, 6", destroyed[198].String())
	assert.Equal(t, "8, 2", destroyed[199].String())
	assert.Equal(t, "10, 9", destroyed[200].String())
	assert.Equal(t, "11, 1", destroyed[298].String())
}
