package main

import (
	"advent-2019/intcode"
	"advent-2019/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	program := utils.ReadCSVNumbers("./day25/input.txt")
	i := intcode.CreateIntcodeComputer(program...)

	input := make(chan string)
	output := i.RunAsciiAsync(input, func() {
		os.Exit(0)
	})

	go func() {
		for x := range output {
			fmt.Print(x)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		input <- text
	}
}
