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

func TestComputerWithAsyncInputAndOutput(t *testing.T) {
	i := CreateIntcodeComputer(3,0,4,0,99)
	input := make(chan int)
	output := i.RunAsync(input)

	input <- 42

	assert.Equal(t, 42, <- output)
}

func TestComputerWithNegativeIntegers(t *testing.T) {
	i := CreateIntcodeComputer(1101,100,-1,4,0)
	i.Run()

	assert.Equal(t, 99, i.Read(4))
}

func TestEqualsUsingPositionMode(t *testing.T) {
	program := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}

	i := CreateIntcodeComputer(program...)
	outputs := i.Run(8)
	assert.Equal(t, 1, outputs[0])

	outputs = i.Run(7)
	assert.Equal(t, 0, outputs[0])
}

func TestLessThanUsingPositionMode(t *testing.T) {
	program := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}

	i := CreateIntcodeComputer(program...)
	outputs := i.Run(8)
	assert.Equal(t, 0, outputs[0])

	outputs = i.Run(7)
	assert.Equal(t, 1, outputs[0])
}

func TestEqualsUsingImmediateMode(t *testing.T) {
	program := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}

	i := CreateIntcodeComputer(program...)
	outputs := i.Run(8)
	assert.Equal(t, 1, outputs[0])

	outputs = i.Run(7)
	assert.Equal(t, 0, outputs[0])
}

func TestLessThanUsingImmediateMode(t *testing.T) {
	program := []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}

	i := CreateIntcodeComputer(program...)
	outputs := i.Run(8)
	assert.Equal(t, 0, outputs[0])

	outputs = i.Run(7)
	assert.Equal(t, 1, outputs[0])
}

func TestJumpingUsingPositionMode(t *testing.T) {
	program := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}

	i := CreateIntcodeComputer(program...)
	outputs := i.Run(0)
	assert.Equal(t, 0, outputs[0])

	outputs = i.Run(42)
	assert.Equal(t, 1, outputs[0])
}

func TestJumpingUsingImmediateMode(t *testing.T) {
	program := []int{3,3,1105,-1,9,1101,0,0,12,4,12,99,1}

	i := CreateIntcodeComputer(program...)
	outputs := i.Run(0)
	assert.Equal(t, 0, outputs[0])

	outputs = i.Run(42)
	assert.Equal(t, 1, outputs[0])
}

func TestJumpingWithAdvancedProgram(t *testing.T) {
	program := []int{
		3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
		1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
		999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99}

	i := CreateIntcodeComputer(program...)
	outputs := i.Run(7)
	assert.Equal(t, 999, outputs[0])

	outputs = i.Run(8)
	assert.Equal(t, 1000, outputs[0])

	outputs = i.Run(9)
	assert.Equal(t, 1001, outputs[0])
}
