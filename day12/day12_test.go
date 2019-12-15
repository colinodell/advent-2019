package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimulateSystem(t *testing.T) {
	input := `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`

	moons := loadMoons(input)

	SimulateSystem(moons, 10)

	assert.Equal(t, "pos=<x=2, y=1, z=-3>, vel=<x=-3, y=-2, z=1>", moons[0].String())
	assert.Equal(t, "pos=<x=1, y=-8, z=0>, vel=<x=-1, y=1, z=3>", moons[1].String())
	assert.Equal(t, "pos=<x=3, y=-6, z=1>, vel=<x=3, y=2, z=-3>", moons[2].String())
	assert.Equal(t, "pos=<x=2, y=0, z=4>, vel=<x=1, y=-1, z=-1>", moons[3].String())

	assert.Equal(t, 179, moons.CalcEnergy())
}

func TestSimulateSystemUntilMatchesInitialState(t *testing.T) {
	input := `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`

	moons := loadMoons(input)

	steps := SimulateSystemUntilMatchesInitialState(moons)

	assert.Equal(t, 4686774924, steps)

}
