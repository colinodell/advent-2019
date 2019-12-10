package intcode

type Intcode struct {
	memory []int
}

func CreateIntcodeComputer(memory []int) Intcode {
	i := Intcode{}
	i.memory = make([]int, len(memory))
	copy(i.memory, memory)

	return i
}

func (i *Intcode) ChangeNounAndVerb(noun, verb int) {
	i.memory[1], i.memory[2] = noun, verb
}

func (i *Intcode) Run() {
	loop: for pos := 0; ; {
		opcode := i.memory[pos]
		switch opcode {
		case 1: // Addition
			operand1 := i.get(i.memory[pos+1])
			operand2 := i.get(i.memory[pos+2])
			i.set(i.memory[pos+3], operand1 + operand2)
			pos += 4
		case 2: // Multiplication
			operand1 := i.get(i.memory[pos+1])
			operand2 := i.get(i.memory[pos+2])
			i.set(i.memory[pos+3], operand1 * operand2)
			pos += 4
		case 99:
			break loop
		}
	}
}

func (i *Intcode) Read(position int) int {
	return i.memory[position]
}

func (i *Intcode) get(arg int) int {
	return i.memory[arg]
}

func (i *Intcode) set(location int, value int) {
	i.memory[location] = value
}
