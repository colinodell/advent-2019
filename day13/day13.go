package main

import (
	"advent-2019/intcode"
	"advent-2019/utils"
	"fmt"
)

const (
	Empty = iota
	Wall
	Block
	Paddle
	Ball
)

func main() {
	g := NewArcadeCabinet()

	fmt.Println("----- Part 1 -----")
	g.Run()
	fmt.Printf("Blocks on screen at game start: %d\n\n", countBlocks(g.screen))

	fmt.Println("----- Part 2 -----")
	g.InsertQuarters(2)
	g.Run()
	fmt.Printf("Final score: %d\n\n", g.score)
}

type Screen map[utils.Vector2]int

type ArcadeCabinet struct {
	computer intcode.Intcode
	screen Screen
	ball, paddle utils.Vector2
	score int
}

func NewArcadeCabinet() *ArcadeCabinet {
	program := utils.ReadCSVNumbers("./day13/input.txt")

	game := new(ArcadeCabinet)
	game.computer = intcode.CreateIntcodeComputer(program...)
	game.screen = make(Screen)

	return game
}

func (g *ArcadeCabinet) InsertQuarters(quarters int) {
	g.computer.OverwriteProgram(0, 2)
}

func (g *ArcadeCabinet) Run() {
	input := make(chan int)
	done := make(chan bool)

	g.score = 0

	output := g.computer.RunAsync(input, func(){ close(done) })

	for {
		select {
		case x := <-output:
			y, tile := <-output, <-output

			if x == -1 && y == 0 {
				g.score = tile
			} else {
				v := utils.Vector2{X: x, Y: y}
				g.screen[v] = tile

				switch tile {
				case Ball:
					g.ball = v
				case Paddle:
					g.paddle = v
				}
			}

		case input <- utils.Sign(g.ball.X - g.paddle.X):
			// Joystick has been moved in the direction of the ball

		case <-done:
			close(input)
			return
		}
	}
}

func countBlocks(screen Screen) int {
	count := 0
	for _, tile := range screen {
		if tile == Block {
			count++
		}
	}
	return count
}
