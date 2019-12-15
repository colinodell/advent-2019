package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 3, Min(3, 42))
	assert.Equal(t, 3, Min(42, 3))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 42, Max(3, 42))
	assert.Equal(t, 42, Max(42, 3))
}
