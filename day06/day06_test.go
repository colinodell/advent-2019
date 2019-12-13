package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const testMap string = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

const partTwoMap string = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`

func TestOrbitCount(t *testing.T) {
	orbits := loadOrbits(strings.Split(testMap, "\n"))
	assert.Equal(t, 42, CountDirectAndIndirectOrbits(orbits))
}

func TestCalculateOrbitalTransfers(t *testing.T) {
	orbits := loadOrbits(strings.Split(partTwoMap, "\n"))
	assert.Equal(t, 4, CalculateOrbitalTransfers(orbits))
}
