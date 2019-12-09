package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateBaseFuelRequirement(t *testing.T) {
	assert.Equal(t, 2, calculateFuelRequirement(12))
	assert.Equal(t, 2, calculateFuelRequirement(14))
	assert.Equal(t, 654, calculateFuelRequirement(1969))
	assert.Equal(t, 33583, calculateFuelRequirement(100756))
}

func TestCalculateTotalFuelRequirement(t *testing.T) {
	assert.Equal(t, 2, calculateTotalFuelRequirement(12))
	assert.Equal(t, 966, calculateTotalFuelRequirement(1969))
	assert.Equal(t, 50346, calculateTotalFuelRequirement(100756))
}