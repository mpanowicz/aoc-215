package main

import (
	"aoc/internal/helpers"
	"bufio"
	"math"
	"os"
	"strings"
)

const (
	Capacity   = "capacity"
	Durability = "durability"
	Flavor     = "flavor"
	Texture    = "texture"
	Calories   = "calories"
)

type NutritionalValue struct {
	Capacity   int
	Durability int
	Flavour    int
	Texture    int
	Calories   int
}

func (n NutritionalValue) Score() int {
	if n.Capacity <= 0 || n.Durability <= 0 || n.Flavour <= 0 || n.Texture <= 0 {
		return 0
	}
	return n.Capacity * n.Durability * n.Flavour * n.Texture
}

func (x NutritionalValue) Add(y NutritionalValue) NutritionalValue {
	return NutritionalValue{
		x.Capacity + y.Capacity,
		x.Durability + y.Durability,
		x.Flavour + y.Flavour,
		x.Texture + y.Texture,
		x.Calories + y.Calories,
	}
}

type Ingredient struct {
	Name string
	NutritionalValue
}

func (i Ingredient) Score(spoons int) NutritionalValue {
	return NutritionalValue{
		Capacity:   i.Capacity * spoons,
		Durability: i.Durability * spoons,
		Flavour:    i.Flavour * spoons,
		Texture:    i.Texture * spoons,
		Calories:   i.Calories * spoons,
	}
}

func getInput() map[string]Ingredient {
	f, _ := os.Open("cmd/day15/input.txt")
	r := bufio.NewReader(f)

	ingredients := map[string]Ingredient{}
	for {
		l, _ := r.ReadString('\n')
		if len(l) == 0 {
			break
		}

		line := l[:len(l)-1]
		nameSplit := strings.Split(line, ": ")
		elements := strings.Split(nameSplit[1], ", ")
		values := map[string]int{}
		for _, e := range elements {
			parts := strings.Split(e, " ")
			values[parts[0]] = helpers.ParseInt(parts[1])
		}
		i := Ingredient{
			nameSplit[0],
			NutritionalValue{
				Capacity:   values[Capacity],
				Durability: values[Durability],
				Flavour:    values[Flavor],
				Texture:    values[Texture],
				Calories:   values[Calories],
			},
		}
		ingredients[nameSplit[0]] = i
	}

	return ingredients
}

func solution() (int, int) {
	MAX_SPOONS := 100
	ingredients := getInput()
	max := math.MinInt
	maxCalories := math.MinInt
	for sprinkles := 0; sprinkles <= MAX_SPOONS; sprinkles++ {
		for butterscotch := 0; butterscotch <= MAX_SPOONS-sprinkles; butterscotch++ {
			for chocolate := 0; chocolate <= MAX_SPOONS-sprinkles-butterscotch; chocolate++ {
				candy := MAX_SPOONS - sprinkles - butterscotch - chocolate
				value := NutritionalValue{}
				value = value.Add(ingredients["Sprinkles"].Score(sprinkles))
				value = value.Add(ingredients["Butterscotch"].Score(butterscotch))
				value = value.Add(ingredients["Chocolate"].Score(chocolate))
				value = value.Add(ingredients["Candy"].Score(candy))
				score := value.Score()
				if score > max {
					max = score
				}
				if value.Calories == 500 && score > maxCalories {
					maxCalories = score
				}
			}
		}
	}

	return max, maxCalories
}

func main() {
	helpers.PrintResult(solution())
}
