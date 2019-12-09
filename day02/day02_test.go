package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputerExecution(t *testing.T) {
	assert.Equal(t, 3500, emulate([]int{1,9,10,3,2,3,11,0,99,30,40,50}, 9, 10))
	assert.Equal(t, 2, emulate([]int{1,0,0,0,99}, 0, 0))
	assert.Equal(t, 2, emulate([]int{2,3,0,3,99}, 3, 0))
	assert.Equal(t, 2, emulate([]int{2,4,4,5,99,0}, 4, 4))
	assert.Equal(t, 30, emulate([]int{1,1,1,4,99,5,6,0,99}, 1, 1))
}
