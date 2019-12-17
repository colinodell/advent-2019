package main

import (
	"advent-2019/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePattern(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1, 0, 0, -1, -1, 0, 0, 1, 1, 0, 0, -1, -1}, generatePattern(15, 2))
	assert.Equal(t, []int{0, 0, 1, 1, 1, 0, 0, 0, -1, -1, -1}, generatePattern(11, 3))
}

func TestIntSliceToString(t *testing.T) {
	assert.Equal(t, "0123456789", intSliceToString([]int{0,1,2,3,4,5,6,7,8,9}))
}

func TestRunFFT(t *testing.T) {
	assert.Equal(t, "12345678", RunFFT("12345678", 0))
	assert.Equal(t, "48226158", RunFFT("12345678", 1))
	assert.Equal(t, "34040438", RunFFT("12345678", 2))
	assert.Equal(t, "03415518", RunFFT("12345678", 3))
	assert.Equal(t, "01029498", RunFFT("12345678", 4))

	assert.Equal(t, "24176176", RunFFT("80871224585914546619083218645595", 100)[:8])
	assert.Equal(t, "73745418", RunFFT("19617804207202209144916044189917", 100)[:8])
	assert.Equal(t, "52432133", RunFFT("69317163492948606335995924319873", 100)[:8])

	input := utils.ReadFile("../day16/input.txt")
	assert.Equal(t, "82525123", RunFFT(input, 100)[:8])
}
