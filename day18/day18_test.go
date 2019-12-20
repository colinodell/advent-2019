package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRunMaze1(t *testing.T) {
	input := strings.Split(`#########
#b.A.@.a#
#########`, "\n")

	maze := NewMaze(input)

	assert.Equal(t, 8, maze.FindShortestPath())
}

func TestRunMaze2(t *testing.T) {
	input := strings.Split(`########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`, "\n")

	maze := NewMaze(input)

	assert.Equal(t, 86, maze.FindShortestPath())
}

func TestRunMaze3(t *testing.T) {
	input := strings.Split(`########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`, "\n")

	maze := NewMaze(input)

	assert.Equal(t, 132, maze.FindShortestPath())
}

func TestRunMaze4(t *testing.T) {
	input := strings.Split(`#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`, "\n")

	maze := NewMaze(input)

	assert.Equal(t, 136, maze.FindShortestPath())
}

func TestRunMaze5(t *testing.T) {
	input := strings.Split(`########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`, "\n")

	maze := NewMaze(input)

	assert.Equal(t, 81, maze.FindShortestPath())
}
