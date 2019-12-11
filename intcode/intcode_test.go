package intcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputerWithDay02Programs(t *testing.T) {
	i := CreateIntcodeComputer(1,9,10,3,2,3,11,0,99,30,40,50)
	i.Run()
	assert.Equal(t, 3500, i.Read(0))

	i.Load(1,0,0,0,99)
	i.Run()
	assert.Equal(t, 2, i.Read(0))

	i.Load(2,3,0,3,99)
	i.Run()
	assert.Equal(t, 2, i.Read(0))

	i.Load(2,4,4,5,99,0)
	i.Run()
	assert.Equal(t, 2, i.Read(0))

	i.Load(1,1,1,4,99,5,6,0,99)
	i.Run()
	assert.Equal(t, 30, i.Read(0))
}

func TestComputerWithParameterModes(t *testing.T) {
	i := CreateIntcodeComputer(1002,4,3,4,33)
	i.Run()
	assert.Equal(t, 99, i.Read(4))
}

func TestComputerWithInputAndOutput(t *testing.T) {
	i := CreateIntcodeComputer(3,0,4,0,99)
	output := i.Run(42)
	assert.Equal(t, 42, output[0])
}

func TestComputerWithNegativeIntegers(t *testing.T) {
	i := CreateIntcodeComputer(1101,100,-1,4,0)
	i.Run()

	assert.Equal(t, 99, i.Read(4))
}
