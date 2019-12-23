package main

import (
	"advent-2019/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input := utils.ReadLines("./day22/input.txt")

	fmt.Println("----- Part 1 -----")
	deck := NewDeck(10007)
	deck.Shuffle(input)
	fmt.Printf("Card 2019 is %d\n\n", deck.FindPosition(2019))
}

type Deck struct {
	Cards []int
}

func NewDeck(n int) *Deck {
	d := new(Deck)
	d.Cards = make([]int, n)
	for i := 0; i < n; i++ {
		d.Cards[i] = i
	}

	return d
}

func (d *Deck) DealIntoNewStack() {
	// Basically just reverse the Cards
	for left, right := 0, len(d.Cards)-1; left < right; left, right = left+1, right-1 {
		d.Cards[left], d.Cards[right] = d.Cards[right], d.Cards[left]
	}
}

func (d *Deck) Cut(n int) {
	if n >= 0 {
		d.Cards = append(d.Cards[n:], d.Cards[:n]...)
	} else {
		d.Cards = append(d.Cards[len(d.Cards)+n:], d.Cards[:len(d.Cards)+n]...)
	}
}

func (d *Deck) DealWithIncrement(n int) {
	cardCount := len(d.Cards)
	newCards := make([]int, cardCount)

	for i := 0; i < cardCount; i++ {
		newCards[(n*i) % cardCount] = d.Cards[i]
	}

	d.Cards = newCards
}

func (d *Deck) Shuffle(instructions []string) {
	regexpCut := regexp.MustCompile(`cut (-?\d+)`)
	regexpIncr := regexp.MustCompile(`deal with increment (\d+)`)

	for _, instruction := range instructions {
		if instruction == "deal into new stack" {
			d.DealIntoNewStack()
		} else if matches := regexpCut.FindStringSubmatch(instruction); matches != nil {
			n, _ := strconv.Atoi(matches[1])
			d.Cut(n)
		} else if matches := regexpIncr.FindStringSubmatch(instruction); matches != nil {
			n, _ := strconv.Atoi(matches[1])
			d.DealWithIncrement(n)
		}
	}
}

func (d *Deck) FindPosition(value int) int {
	for p, v := range d.Cards {
		if v == value {
			return p
		}
	}

	panic("Card not found!")
}
