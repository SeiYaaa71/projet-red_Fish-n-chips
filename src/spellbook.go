package main

import "fmt"

func spellBook(c *Character) {
	sort := "Boule de feu"

	for _, skill := range c.Skill {
		if skill == sort {
			fmt.Println("❌ Vous connaissez déjà ce sort :", sort)
			return
		}
	}

	c.Skill = append(c.Skill, sort)
	fmt.Println("✨ Nouveau sort appris :", sort)
}
