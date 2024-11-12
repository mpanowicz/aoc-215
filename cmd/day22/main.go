package main

import (
	"aoc/internal/helpers"
	"sort"
)

const (
	ShieldArmor  = 7
	PoisonDamage = 3
	ManaRestore  = 101

	MagicMissile = "MagicMissile"
	Drain        = "Drain"
	Shield       = "Shield"
	Poison       = "Poison"
	Recharge     = "Recharge"
)

type Boss struct {
	HitPoints int
	Damage    int

	Poisoned int
}

type Mage struct {
	HitPoints int
	Mana      int

	Shielded  int
	Recharged int

	UsedMana int
}

type Spell func(f Fight)
type SpellAction struct {
	Mana   int
	Action Spell
}

func (m *Mage) UseMana(v int) {
	m.Mana -= v
	m.UsedMana += v
}

func (m *Mage) MagicMissile(b *Boss) {
	m.UseMana(53)
	b.HitPoints -= 4
}

func (m *Mage) Drain(b *Boss) {
	m.UseMana(73)
	m.HitPoints += 2
	b.HitPoints -= 2
}

func (m *Mage) Shield(b *Boss) {
	m.UseMana(113)
	m.Shielded += 6
}

func (m *Mage) Poison(b *Boss) {
	m.UseMana(173)
	b.Poisoned += 6
}

func (m *Mage) Recharge(b *Boss) {
	m.UseMana(229)
	m.Recharged += 5
}

func (m *Mage) TakeDamage(b *Boss) {
	damage := b.Damage
	if m.Shielded > 0 {
		damage -= ShieldArmor
	}
	m.HitPoints -= damage
}

var Spells = map[string]SpellAction{
	MagicMissile: {53, func(f Fight) {
		*f.Actions = append(*f.Actions, MagicMissile)
		f.Mage.MagicMissile(f.Boss)
	}},
	Drain: {73, func(f Fight) {
		*f.Actions = append(*f.Actions, Drain)
		f.Mage.Drain(f.Boss)
	}},
	Shield: {113, func(f Fight) {
		*f.Actions = append(*f.Actions, Shield)
		f.Mage.Shield(f.Boss)
	}},
	Poison: {173, func(f Fight) {
		*f.Actions = append(*f.Actions, Poison)
		f.Mage.Poison(f.Boss)
	}},
	Recharge: {229, func(f Fight) {
		*f.Actions = append(*f.Actions, Recharge)
		f.Mage.Recharge(f.Boss)
	}},
}

type Fight struct {
	Boss *Boss
	Mage *Mage

	Actions *[]string
}

func (f Fight) NextPossibleSpells() []SpellAction {
	spells := []SpellAction{}

	spell := Spells[MagicMissile]
	if spell.Mana <= f.Mage.Mana {
		spells = append(spells, spell)
	}

	spell = Spells[Drain]
	if spell.Mana <= f.Mage.Mana {
		spells = append(spells, spell)
	}

	spell = Spells[Poison]
	if spell.Mana <= f.Mage.Mana && f.Boss.Poisoned == 0 {
		spells = append(spells, spell)
	}

	spell = Spells[Shield]
	if spell.Mana <= f.Mage.Mana && f.Mage.Shielded == 0 {
		spells = append(spells, spell)
	}

	spell = Spells[Recharge]
	if spell.Mana <= f.Mage.Mana && f.Mage.Recharged == 0 {
		spells = append(spells, spell)
	}

	return spells
}

func (f Fight) NextTurn(losePoint int) []Fight {
	fights := []Fight{}
	f.Mage.HitPoints -= losePoint

	if f.Mage.Shielded > 0 {
		f.Mage.Shielded--
	}

	if f.Mage.Recharged > 0 {
		f.Mage.Mana += ManaRestore
		f.Mage.Recharged--
	}

	if f.Boss.Poisoned > 0 {
		f.Boss.HitPoints -= PoisonDamage
		f.Boss.Poisoned--
	}

	if f.Boss.HitPoints <= 0 {
		return []Fight{f}
	}

	actions := f.NextPossibleSpells()
	if len(actions) == 0 {
		return fights
	}

	for _, a := range actions {
		bossCopy := Boss(*f.Boss)
		mageCopy := Mage(*f.Mage)
		actionsCopy := make([]string, len(*f.Actions))
		copy(actionsCopy, *f.Actions)
		c := Fight{&bossCopy, &mageCopy, &actionsCopy}
		a.Action(c)

		if c.Mage.Shielded > 0 {
			c.Mage.Shielded--
		}

		if c.Mage.Recharged > 0 {
			c.Mage.Mana += ManaRestore
			c.Mage.Recharged--
		}

		if c.Boss.Poisoned > 0 {
			c.Boss.HitPoints -= PoisonDamage
			c.Boss.Poisoned--
		}

		if c.Boss.HitPoints <= 0 {
			fights = append(fights, c)
		} else {
			c.Mage.TakeDamage(c.Boss)

			fights = append(fights, c)
		}
	}
	return fights
}

func (f Fight) Print() struct {
	Boss    Boss
	Mage    Mage
	Actions []string
} {
	return struct {
		Boss    Boss
		Mage    Mage
		Actions []string
	}{*f.Boss, *f.Mage, *f.Actions}
}

func simulate(fights []Fight, losePoint int) int {
	for {
		nextFights := []Fight{}
		for _, f := range fights {
			if f.Boss.HitPoints <= 0 && f.Mage.HitPoints > 0 {
				return f.Mage.UsedMana
			}

			if f.Mage.HitPoints <= 0 {
				continue
			}

			next := f.NextTurn(losePoint)
			nextFights = append(nextFights, next...)
		}
		sort.Slice(nextFights, func(i, j int) bool { return nextFights[i].Mage.UsedMana < nextFights[j].Mage.UsedMana })
		fights = nextFights
	}
}

func solution() (int, int) {
	boss := Boss{HitPoints: 51, Damage: 9}
	mage := Mage{HitPoints: 50, Mana: 500, UsedMana: 0}
	fights := []Fight{{&boss, &mage, &[]string{}}}
	p1 := simulate(fights, 0)
	p2 := simulate(fights, 1)

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
