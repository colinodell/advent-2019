package intcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseOpcode(t *testing.T) {
	opcode, paramModes := parseOpcode(2)
	assert.Equal(t, 2, opcode)
	assert.Equal(t, []int{0,0,0}, paramModes)

	opcode, paramModes = parseOpcode(21005)
	assert.Equal(t, 5, opcode)
	assert.Equal(t, []int{0,1,2}, paramModes)
}

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
	output := i.RunAsync(input, nil)

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

func TestRelativeBaseReadFunctionality(t *testing.T) {
	// This program takes no input and produces a copy of itself
	program1 := []int{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99}
	i := CreateIntcodeComputer(program1...)
	outputs := i.Run()
	assert.Equal(t, program1, outputs)

	// This program outputs a 16-digit number
	i.Load(1102,34915192,34915192,7,4,7,99,0)
	outputs = i.Run()
	assert.Equal(t, 16, countDigits(outputs[0]))

	// This program will output the large number in the middle
	i.Load(104,1125899906842624,99)
	outputs = i.Run()
	assert.Equal(t, 1125899906842624, outputs[0])

	// Custom program to test relative mode opcode
	i.Load(109, 2000, 109, 19, 109, -34, 99)
	i.Run()
	assert.Equal(t, 1985, i.relativeBase)
}

func TestRelativeBaseWriteFunctionality(t *testing.T) {
	// Test programs from https://www.reddit.com/r/adventofcode/comments/e8aw9j/2019_day_9_part_1_how_to_fix_203_error/fac3294/
	i := CreateIntcodeComputer()

	i.Load(109, -1, 4, 1, 99)
	outputs := i.Run()
	assert.Equal(t, -1, outputs[0])

	i.Load(109, -1, 104, 1, 99)
	outputs = i.Run()
	assert.Equal(t, 1, outputs[0])

	i.Load(109, -1, 204, 1, 99)
	outputs = i.Run()
	assert.Equal(t, 109, outputs[0])

	i.Load(109, 1, 9, 2, 204, -6, 99)
	outputs = i.Run()
	assert.Equal(t, 204, outputs[0])

	i.Load(109, 1, 109, 9, 204, -6, 99)
	outputs = i.Run()
	assert.Equal(t, 204, outputs[0])

	i.Load(109, 1, 209, -1, 204, -106, 99)
	outputs = i.Run()
	assert.Equal(t, 204, outputs[0])

	i.Load(109, 1, 3, 3, 204, 2, 99)
	outputs = i.Run(42)
	assert.Equal(t, 42, outputs[0])

	i.Load(109, 1, 203, 2, 204, 2, 99)
	outputs = i.Run(42)
	assert.Equal(t, 42, outputs[0])
}

func countDigits(i int) (count int) {
	for i != 0 {

		i /= 10
		count = count + 1
	}
	return count
}
