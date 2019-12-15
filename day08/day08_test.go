package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImage_GetLayerWithFewest(t *testing.T) {
	image := NewImage("123456789012", 3, 2)
	layer := image.GetLayerWithFewest(0)
	assert.Equal(t, []int{1,2,3,4,5,6}, layer.data)
}

func TestImage_Render(t *testing.T) {
	image := NewImage("0222112222120000", 2, 2)
	assert.Equal(t, " X\nX \n", image.Render())
}
