package main

import "fmt"

func spellBook(c *Character) {
	sort := "Boule de feu"

	for _, skill := range c.Skills {
		if skill == sort {
			fmt.Println("❌ Vous connaissez déjà ce sort :", sort)
			return
		}
	}

	c.Skills = append(c.Skills, sort)
	fmt.Println("✨ Nouveau sort appris :", sort)
}
