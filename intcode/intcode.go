package intcode

import (
	"errors"
	"strconv"
	"strings"
)

const (
	param_mode_position = 0
	param_mode_immediate = 1
	param_mode_relative = 2

	max_parameters = 3

	max_memory = 1024*1024 // 1 MB
)

type Intcode struct {
	program                  []int
	memory                   []int
	pos                      int
	input                    chan int
	output                   chan int
	relativeBase             int
}

func CreateIntcodeComputer(program ...int) *Intcode {
	i := Intcode{}
	i.Load(program...)

	return &i
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
	for outputValue := range i.RunAsync(inputChannel, nil) {
		outputs = append(outputs, outputValue)
	}

	return outputs
}

func (i *Intcode) RunAscii(program string) (int, string) {
	input := make(chan int)
	done := make(chan bool)
	output := i.RunAsync(input, func(){ done <- true })

	go func() {
		for _, char := range program {
			input <- int(char)
		}
		close(input)
	}()

	var result int
	var display strings.Builder

	for {
		select {
		case x := <-output:
			if x >= 128 {
				result = x
			} else {
				display.WriteRune(rune(x))
			}
		case <-done:
			return result, display.String()
		}
	}
}

func (i *Intcode) RunAsync(input chan int, done func()) chan int {

	i.pos = 0
	i.relativeBase = 0
	i.memory = make([]int, max_memory)
	copy(i.memory, i.program)

	i.input = input
	i.output = make(chan int)

	go func() {
		for {
			err := i.executeNextOperation()

			if err != nil {
				close(i.output)
				if done != nil {
					done()
				}
				return
			}
		}
	}()

	return i.output
}

func (i *Intcode) SetOutput(output chan int) {
	close(i.output)
	i.output = output
}

func (i *Intcode) executeNextOperation() error {
	if i.pos >= len(i.memory) {
		panic("oops")
	}
	opcode, parameterModes := parseOpcode(i.memory[i.pos])
	switch opcode {
	case 1: // Addition
		operand1 := i.get(i.memory[i.pos+1], parameterModes[0])
		operand2 := i.get(i.memory[i.pos+2], parameterModes[1])
		i.set(i.memory[i.pos+3], operand1 + operand2, parameterModes[2])
		i.pos += 4
	case 2: // Multiplication
		operand1 := i.get(i.memory[i.pos+1], parameterModes[0])
		operand2 := i.get(i.memory[i.pos+2], parameterModes[1])
		i.set(i.memory[i.pos+3], operand1 * operand2, parameterModes[2])
		i.pos += 4
	case 3: // input
		i.set(i.memory[i.pos+1], <-i.input, parameterModes[0])
		i.pos += 2
	case 4: // output
		i.output <- i.get(i.memory[i.pos+1], parameterModes[0])
		i.pos += 2
	case 5: // Jump If True
		operand1 := i.get(i.memory[i.pos+1], parameterModes[0])
		operand2 := i.get(i.memory[i.pos+2], parameterModes[1])

		if operand1 != 0 {
			i.pos = operand2
		} else {
			i.pos += 3
		}
	case 6: // Jump If False
		operand1 := i.get(i.memory[i.pos+1], parameterModes[0])
		operand2 := i.get(i.memory[i.pos+2], parameterModes[1])

		if operand1 == 0 {
			i.pos = operand2
		} else {
			i.pos += 3
		}
	case 7: // Less Than
		operand1 := i.get(i.memory[i.pos+1], parameterModes[0])
		operand2 := i.get(i.memory[i.pos+2], parameterModes[1])
		if operand1 < operand2 {
			i.set(i.memory[i.pos+3], 1, parameterModes[2])
		} else {
			i.set(i.memory[i.pos+3], 0, parameterModes[2])
		}
		i.pos += 4
	case 8: // Equals
		operand1 := i.get(i.memory[i.pos+1], parameterModes[0])
		operand2 := i.get(i.memory[i.pos+2], parameterModes[1])
		if operand1 == operand2 {
			i.set(i.memory[i.pos+3], 1, parameterModes[2])
		} else {
			i.set(i.memory[i.pos+3], 0, parameterModes[2])
		}
		i.pos += 4
	case 9: // Adjust relative base
		i.relativeBase += i.get(i.memory[i.pos+1], parameterModes[0])
		i.pos += 2
	case 99:
		return errors.New("execution halted")
	default:
		panic("Unknown opcode")
	}

	return nil
}

func (i *Intcode) Read(position int) int {
	return i.memory[position]
}

func (i *Intcode) get(arg int, mode int) int {
	switch mode {
	case param_mode_position:
		return i.memory[arg]
	case param_mode_immediate:
		return arg
	case param_mode_relative:
		return i.memory[i.relativeBase + arg]
	}

	panic("Invalid param mode")
}

func (i *Intcode) set(location int, value int, mode int) {
	switch mode {
	case param_mode_position:
		i.memory[location] = value
	case param_mode_relative:
		i.memory[i.relativeBase + location] = value
	}
}

func (i *Intcode) OverwriteProgram(location int, value int) {
	i.program[location] = value
}

func parseOpcode(instruction int) (int, []int) {
	opcode := instruction % 100

	paramModes := make([]int, max_parameters)
	paramModesAsString := strconv.Itoa(10000000 + ((instruction - opcode) / 100))
	for i := 0; i < max_parameters; i++ {
		paramModes[i], _ = strconv.Atoi(string(paramModesAsString[len(paramModesAsString) - i - 1]))
	}

	return opcode, paramModes
}
