package main

import "fmt"

func canAddItem(c *Character) bool {
	if len(c.Inventaire) >= 10 {
		fmt.Println(Red + "⚠️ Inventaire plein, impossible d'ajouter cet item !" + Reset)
		waitForEnter()
		return false
	}
	return true
}
