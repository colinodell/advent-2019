package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck(10)

	assert.Equal(t, []int{0,1,2,3,4,5,6,7,8,9}, deck.Cards)
}

func TestDeck_DealIntoNewStack(t *testing.T) {
	deck := NewDeck(10)
	deck.DealIntoNewStack()

	assert.Equal(t, []int{9,8,7,6,5,4,3,2,1,0}, deck.Cards)
}

func TestDeck_CutPositive(t *testing.T) {
	deck := NewDeck(10)
	deck.Cut(3)

	assert.Equal(t, []int{3,4,5,6,7,8,9,0,1,2}, deck.Cards)
}

func TestDeck_CutNegative(t *testing.T) {
	deck := NewDeck(10)
	deck.Cut(-4)

	assert.Equal(t, []int{6,7,8,9,0,1,2,3,4,5}, deck.Cards)
}

func TestDeck_DealWithIncrement(t *testing.T) {
	deck := NewDeck(10)
	deck.DealWithIncrement(3)

	assert.Equal(t, []int{0,7,4,1,8,5,2,9,6,3}, deck.Cards)
}

func TestDeck_Shuffle(t *testing.T) {
	instructions := strings.Split(`deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1`, "\n")

	deck := NewDeck(10)
	deck.Shuffle(instructions)

	assert.Equal(t, []int{9,2,5,8,1,4,7,0,3,6}, deck.Cards)
}
