package main

import (
	"advent-2019/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadLines("./day14/input.txt")
	reactions := loadReactions(input)

	fmt.Println("----- Part 1 -----")
	fmt.Printf("To make 1 FUEL you will need %d ORE\n\n", reactions.CalculateOreRequiredToMakeFuel(1))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("1 trillion ORE can make %d FUEL\n\n", reactions.CalculateMaximumFuel(1_000_000_000_000))
}

type Chemical string

type Ingredient struct {
	Chemical Chemical
	Quantity int
}

type Reaction struct {
	Inputs []Ingredient
	Output Ingredient
}

type Reactions map[Chemical]Reaction

type Inventory map[Chemical]int

func loadReactions(input []string) Reactions {
	reactions := make(Reactions)

	re := regexp.MustCompile(`^(\d+ [A-Z]+(?:, \d+ [A-Z]+)*) => (\d+ [A-Z]+)$`)

	for _, line := range input {
		var reaction Reaction

		matches := re.FindStringSubmatch(line)
		reaction.Output = parseIngredient(matches[2])
		for _, ingredient := range strings.Split(matches[1], ", ") {
			reaction.Inputs = append(reaction.Inputs, parseIngredient(ingredient))
		}

		reactions[reaction.Output.Chemical] = reaction
	}

	return reactions
}

func parseIngredient(s string) Ingredient {
	x := strings.Split(s, " ")

	c := Chemical(x[1])
	q, _ := strconv.Atoi(x[0])

	return Ingredient{
		Chemical: c,
		Quantity: q,
	}
}

func (r Reactions) CalculateOreRequiredToMakeFuel(qty int) int {
	have := make(Inventory)
	want := Ingredient{Chemical: "FUEL", Quantity: qty}

	ore, _ := r.Resolve(have, want)

	return ore
}

func (r Reactions) CalculateMaximumFuel(ore int) int {
	// Perform a binary search with a maximum that assumes 1 ore = 1 fuel
	return sort.Search(ore, func(fuel int) bool {
		return r.CalculateOreRequiredToMakeFuel(fuel) > ore
	}) - 1
}

func (r Reactions) Resolve(have Inventory, want Ingredient) (int, Inventory) {
	if want.Chemical == "ORE" {
		return want.Quantity, have
	}

	// Do we have any leftovers available?
	available := utils.Min(want.Quantity, have[want.Chemical])
	// Use them up
	want.Quantity -= available
	have[want.Chemical] -= available

	// How much more do we need to make?
	reaction := r[want.Chemical]
	reactionsToRun := (want.Quantity + reaction.Output.Quantity - 1) / reaction.Output.Quantity

	// Any leftover quantity should go back into inventory
	have[want.Chemical] += reactionsToRun * reaction.Output.Quantity - want.Quantity

	// Consume the necessary inputs (recursively) to find out how much ore we'll need
	ore := 0
	for _, input := range reaction.Inputs {
		// We'll need a batch of this ingredient for each reaction
		input.Quantity *= reactionsToRun

		// Recursively resolve the components for this ingredient
		var additionalOre int
		additionalOre, have = r.Resolve(have, input)
		ore += additionalOre
	}

	return ore, have
}
