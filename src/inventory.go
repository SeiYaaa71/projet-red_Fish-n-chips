package main

import "fmt"

func addInventory(c *Character, item string) {
	c.Inventaire = append(c.Inventaire, item)
	fmt.Printf("✅ Vous avez obtenu : %s\n", item)
}

func removeInventory(c *Character, index int) {
	if index >= 0 && index < len(c.Inventaire) {
		fmt.Printf("❌ %s a été retiré de l'inventaire.\n", c.Inventaire[index])
		c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)
	} else {
		fmt.Println("Indice invalide, impossible de retirer l'item.")
	}
}
