package main

import (
	"advent-2019/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepairDroidWithPuzzleInput(t *testing.T) {
	program := utils.ReadCSVNumbers("../day15/input.txt")
	droid := NewRepairDroid(program)

	assert.Equal(t, 266, droid.TraverseMaze(true))

	droid.ResetMaze()

	assert.Equal(t, 274, droid.TraverseMaze(false))
}
