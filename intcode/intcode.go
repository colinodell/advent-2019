package intcode

import "strconv"

type Intcode struct {
	memory []int
}

func CreateIntcodeComputer(program ...int) Intcode {
	i := Intcode{}
	i.Load(program...)

	return i
}

func (i *Intcode) Load(program ...int) {
	i.memory = make([]int, len(program))
	copy(i.memory, program)
}

func (i *Intcode) ChangeNounAndVerb(noun, verb int) {
	i.memory[1], i.memory[2] = noun, verb
}

func (i *Intcode) Run(inputs ...int) []int {
	var nextInput int
	var outputs []int

	loop: for pos := 0; ; {
		opcode, parameterModes := parseOpcode(i.memory[pos])
		switch opcode {
		case 1: // Addition
			operand1 := i.get(i.memory[pos+1], parameterModes & 1 != 0)
			operand2 := i.get(i.memory[pos+2], parameterModes & 2 != 0)
			i.set(i.memory[pos+3], operand1 + operand2)
			pos += 4
		case 2: // Multiplication
			operand1 := i.get(i.memory[pos+1], parameterModes & 1 != 0)
			operand2 := i.get(i.memory[pos+2], parameterModes & 2 != 0)
			i.set(i.memory[pos+3], operand1 * operand2)
			pos += 4
		case 3: // Input
			// Grab the next input and remove it
			nextInput, inputs = inputs[0], inputs[1:]
			i.set(i.memory[pos+1], nextInput)
			pos += 2
		case 4: // Output
			outputs = append(outputs, i.get(i.memory[pos+1], false))
			pos += 2

		case 99:
			break loop
		}
	}

	return outputs
}

func (i *Intcode) Read(position int) int {
	return i.memory[position]
}

func (i *Intcode) get(arg int, immediateMode bool) int {
	if immediateMode {
		return arg
	}

	return i.memory[arg]
}

func (i *Intcode) set(location int, value int) {
	i.memory[location] = value
}

func parseOpcode(instruction int) (int, int) {
	opcode := instruction % 100

	paramModesDecimal := (instruction - opcode) / 100

	parameterModes, err := strconv.ParseInt(strconv.Itoa(paramModesDecimal), 2, 64)
	if err != nil {
		panic("Invalid parameter modes")
	}

	return opcode, int(parameterModes)
}
