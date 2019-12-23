package main

import (
	"advent-2019/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	input := utils.ReadCSVNumbers("../day23/input.txt")
	first, last := Run(input)

	assert.Equal(t, 21664, first)
	assert.Equal(t, 16150, last)
}
