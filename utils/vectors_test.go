package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVector2_Add(t *testing.T) {
	v1 := Vector2{X: 1, Y: 2}
	v2 := Vector2{X: 3, Y: 4}

	v3 := v1.Add(v2)

	assert.Equal(t, Vector2{X: 4, Y: 6}, v3)
}

func TestVector2_Min(t *testing.T) {
	v1 := Vector2{X: 9, Y: 1}
	v2 := Vector2{X: 2, Y: 8}

	v3 := v1.Min(v2)

	assert.Equal(t, Vector2{X: 2, Y: 1}, v3)
}

func TestVector2_Max(t *testing.T) {
	v1 := Vector2{X: 9, Y: 1}
	v2 := Vector2{X: 2, Y: 8}

	v3 := v1.Max(v2)

	assert.Equal(t, Vector2{X: 9, Y: 8}, v3)
}
