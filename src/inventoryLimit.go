package main

import "fmt"

func canAddItem(c *Character) bool {
	if len(c.Inventaire) >= 10 {
		fmt.Println("⚠️ Inventaire plein ! Vous ne pouvez pas porter plus de 10 objets.")
		return false
	}
	return true
}
