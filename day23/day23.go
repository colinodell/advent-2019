package main

import (
	"advent-2019/intcode"
	"advent-2019/utils"
	"fmt"
)

func main() {
	input := utils.ReadCSVNumbers("./day23/input.txt")
	first, last := Run(input)
	fmt.Println("----- Part 1 -----")
	fmt.Printf("First value: %d\n\n", first)

	fmt.Println("----- Part 2 -----")

	fmt.Printf("Last value: %d\n", last)
}

func Run(program []int) (int, int) {
	n := NewNetwork(50, program)
	results := n.Run()

	first := <- results

	var last int
	for v := range results {
		last = v
	}

	return first, last
}

type Packet utils.Vector2

type Network struct {
	size int
	in map[int]chan int
	out map[int]chan int
}

func NewNetwork(size int, program []int) *Network {
	n := new(Network)
	n.size = size
	n.in = make(map[int]chan int)
	n.out = make(map[int]chan int)

	for i := 0; i < size; i++ {
		n.in[i] = make(chan int)

		c := intcode.CreateIntcodeComputer(program...)
		n.out[i] = c.RunAsync(n.in[i], nil)

		// Tell the computer its network id
		n.in[i] <- i
		n.in[i] <- -1
	}

	return n
}

func (n *Network) Run() chan int {
	var nat Packet
	var lastNat Packet
	var idle int
	output := make(chan int)

	go func() {
		for i := 0; ; i = (i+1) % n.size {
			select {
			case addr := <-n.out[i]:
				if addr == 255 {
					nat = Packet{X: <-n.out[i], Y: <-n.out[i]}
				} else {
					n.in[addr] <- <-n.out[i]
					n.in[addr] <- <-n.out[i]
				}
				idle = 0

			case n.in[i] <- -1:
				idle++
			}

			if idle >= n.size {
				output <- nat.Y
				if nat.Y == lastNat.Y && nat.Y != 0 {
					close(output)
					return
				}

				n.in[0] <- nat.X
				n.in[0] <- nat.Y
				lastNat = nat
				idle = 0
			}
		}
	}()

	return output
}
