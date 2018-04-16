package main

import "fmt"

type character struct {
	name     string
	level    int
	equipped string
	weapon
	armor
}

type weapon struct {
	axe
	sword
	hammer
}

type axe struct {
	name   string
	damage int
}

type sword struct {
	name   string
	damage int
}

type hammer struct {
	name   string
	damage int
}
type armor struct {
	name    string
	defense int
}

func (att character) att(target string) string {
	var weap string
	var damage int
	equipped := att.equipped
	switch equipped {
	case ("axe"):
		weap = att.weapon.axe.name
		damage = att.weapon.axe.damage
	case ("sword"):
		weap = att.weapon.sword.name
		damage = att.weapon.sword.damage
	case ("hammer"):
		weap = att.weapon.hammer.name
		damage = att.weapon.hammer.damage
	default:
		weap = "hands"
		damage = 1
	}
	fmt.Println(att.name, "attacks", target, "with", weap, "for", damage)
	return att.name
}

func main() {
	c1 := character{
		name:     "Tim",
		level:    1,
		equipped: "axe",
		weapon: weapon{
			axe: axe{
				name:   "Gloom",
				damage: 10,
			},
		},
		armor: armor{
			name:    "Robe",
			defense: 3,
		},
	}

	fmt.Println(c1)

	c2 := character{
		name:     "Robin",
		level:    1,
		equipped: "sword",
		weapon: weapon{
			sword: sword{
				name:   "Sting",
				damage: 5,
			},
		},
		armor: armor{
			name:    "Breastplate",
			defense: 5,
		},
	}

	fmt.Println(c2)
	fmt.Println(c1.att(c2.name))
}
