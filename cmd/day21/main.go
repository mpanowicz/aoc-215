package main

import (
	"aoc/internal/helpers"
	"fmt"
	"slices"
	"sort"
)

type Warrior struct {
	HitPoints int

	Build Build
}

type Build struct {
	Weapon *Item
	Armor  *Item
	Ring1  *Item
	Ring2  *Item
}

func (b Build) Damage() int {
	damage := b.Weapon.Damage
	if b.Ring1 != nil {
		damage += b.Ring1.Damage
	}
	if b.Ring2 != nil {
		damage += b.Ring2.Damage
	}
	return damage
}

func (b Build) Defense() int {
	defense := 0
	if b.Armor != nil {
		defense += b.Armor.Armor
	}
	if b.Ring1 != nil {
		defense += b.Ring1.Armor
	}
	if b.Ring2 != nil {
		defense += b.Ring2.Armor
	}
	return defense
}

func (b Build) Cost() int {
	cost := b.Weapon.Cost
	if b.Armor != nil {
		cost += b.Armor.Cost
	}
	if b.Ring1 != nil {
		cost += b.Ring1.Cost
	}
	if b.Ring2 != nil {
		cost += b.Ring2.Cost
	}
	return cost
}

func (b Build) Print() string {
	w := Item{}
	a := Item{}
	r1 := Item{}
	r2 := Item{}
	if b.Weapon != nil {
		w = *b.Weapon
	}
	if b.Armor != nil {
		a = *b.Armor
	}
	if b.Ring1 != nil {
		r1 = *b.Ring1
	}
	if b.Ring2 != nil {
		r2 = *b.Ring2
	}
	return fmt.Sprintf("W:%v A:%v R1:%v R2:%v Cost:%d", w, a, r1, r2, b.Cost())
}

type Item struct {
	Cost   int
	Damage int
	Armor  int
}

var weapons = []Item{
	{8, 4, 0},
	{10, 5, 0},
	{25, 6, 0},
	{40, 7, 0},
	{74, 8, 0},
}

var armors = []Item{
	{13, 0, 1},
	{31, 0, 2},
	{53, 0, 3},
	{75, 0, 4},
	{102, 0, 5},
}

var rings = []Item{
	{20, 0, 1},
	{25, 1, 0},
	{40, 0, 2},
	{50, 2, 0},
	{80, 0, 3},
	{100, 3, 0},
}

func getBuilds() []Build {
	builds := []Build{}
	for _, w := range weapons {
		build := Build{Weapon: &w}
		builds = append(builds, build)

		for _, a := range armors {
			build.Armor = &a
			builds = append(builds, build)

			for _, r1 := range rings {
				build.Ring1 = &r1
				build.Ring2 = nil
				builds = append(builds, build)
				for _, r2 := range rings {
					if r2 != r1 {
						build.Ring2 = &r2
						builds = append(builds, build)
					}
				}
			}
		}

		build.Armor = nil
		build.Ring1 = nil
		build.Ring2 = nil
		for i := 0; i < len(rings); i++ {
			build.Ring1 = &rings[i]
			build.Ring2 = nil
			builds = append(builds, build)
			for j := 0; j < len(rings); j++ {
				if i != j {
					build.Ring2 = &rings[j]
					builds = append(builds, build)
				}
			}
		}
	}
	return builds
}

func won(you Warrior, boss Warrior) bool {
	turn := 1
	for {
		if boss.HitPoints <= 0 {
			return true
		} else if you.HitPoints <= 0 {
			return false
		}

		if turn == 1 {
			boss.HitPoints -= (you.Build.Damage() - boss.Build.Defense())
		} else {
			you.HitPoints -= (boss.Build.Damage() - you.Build.Defense())
		}

		turn *= -1
	}
}

func solution() (int, int) {
	boss := Warrior{HitPoints: 109, Build: Build{Weapon: &Item{0, 8, 0}, Armor: &Item{0, 0, 2}}}

	me := Warrior{HitPoints: 100}

	builds := getBuilds()
	sort.Slice(builds, func(a, b int) bool { return builds[a].Cost() < builds[b].Cost() })
	for _, b := range builds {
		me.Build = b
		if won(me, boss) {
			break
		}
	}
	p1 := me.Build

	slices.Reverse(builds)
	for _, b := range builds {
		me.Build = b
		if !won(me, boss) {
			break
		}
	}
	p2 := me.Build

	return p1.Cost(), p2.Cost()
}

func main() {
	helpers.PrintResult(solution())
}
