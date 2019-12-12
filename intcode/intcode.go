package intcode

import (
	"errors"
	"strconv"
)

type Intcode struct {
	program []int
	memory  []int
	pos     int
	input   chan int
	output  chan int
}

func CreateIntcodeComputer(program ...int) Intcode {
	i := Intcode{}
	i.Load(program...)

	return i
}

func (i *Intcode) Load(program ...int) {
	i.program = make([]int, len(program))
	copy(i.program, program)
}

func (i *Intcode) ChangeNounAndVerb(noun, verb int) {
	i.program[1], i.program[2] = noun, verb
}

func (i *Intcode) Run(inputs ...int) []int {
	inputChannel := make(chan int)

	// Pump all the inputs into the channel
	go func() {
		for _, value := range inputs {
			inputChannel <- value
		}
		close(inputChannel)
	}()

	// Create the output channel
	var outputs []int

	// Run the program and read outputs into our slice
	for outputValue := range i.RunAsync(inputChannel) {
		outputs = append(outputs, outputValue)
	}

	return outputs
}

func (i *Intcode) RunAsync(input chan int) chan int {
	i.pos = 0
	i.memory = make([]int, len(i.program))
	copy(i.memory, i.program)

	i.input = input
	i.output = make(chan int)

	go func() {
		for {
			err := i.executeNextOperation()

			if err != nil {
				close(i.output)
				return
			}
		}
	}()

	return i.output
}

func (i *Intcode) executeNextOperation() error {
	if i.pos >= len(i.memory) {
		panic("oops")
	}
	opcode, parameterModes := parseOpcode(i.memory[i.pos])
	switch opcode {
	case 1: // Addition
		operand1 := i.get(i.memory[i.pos+1], parameterModes & 1 != 0)
		operand2 := i.get(i.memory[i.pos+2], parameterModes & 2 != 0)
		i.set(i.memory[i.pos+3], operand1 + operand2)
		i.pos += 4
	case 2: // Multiplication
		operand1 := i.get(i.memory[i.pos+1], parameterModes & 1 != 0)
		operand2 := i.get(i.memory[i.pos+2], parameterModes & 2 != 0)
		i.set(i.memory[i.pos+3], operand1 * operand2)
		i.pos += 4
	case 3: // input
		i.set(i.memory[i.pos+1], <-i.input)
		i.pos += 2
	case 4: // output
		i.output <- i.get(i.memory[i.pos+1], parameterModes & 1 != 0)
		i.pos += 2
	case 5: // Jump If True
		operand1 := i.get(i.memory[i.pos+1], parameterModes & 1 != 0)
		operand2 := i.get(i.memory[i.pos+2], parameterModes & 2 != 0)

		if operand1 != 0 {
			i.pos = operand2
		} else {
			i.pos += 3
		}
	case 6: // Jump If False
		operand1 := i.get(i.memory[i.pos+1], parameterModes & 1 != 0)
		operand2 := i.get(i.memory[i.pos+2], parameterModes & 2 != 0)

		if operand1 == 0 {
			i.pos = operand2
		} else {
			i.pos += 3
		}
	case 7: // Less Than
		operand1 := i.get(i.memory[i.pos+1], parameterModes & 1 != 0)
		operand2 := i.get(i.memory[i.pos+2], parameterModes & 2 != 0)
		if operand1 < operand2 {
			i.set(i.memory[i.pos+3], 1)
		} else {
			i.set(i.memory[i.pos+3], 0)
		}
		i.pos += 4
	case 8: // Equals
		operand1 := i.get(i.memory[i.pos+1], parameterModes & 1 != 0)
		operand2 := i.get(i.memory[i.pos+2], parameterModes & 2 != 0)
		if operand1 == operand2 {
			i.set(i.memory[i.pos+3], 1)
		} else {
			i.set(i.memory[i.pos+3], 0)
		}
		i.pos += 4
	case 99:
		return errors.New("execution halted")
	}

	return nil
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
