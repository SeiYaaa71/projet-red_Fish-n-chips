package main

import "fmt"

// Fonction pour utiliser une potion
func takePot(c *Character) {
	// Vérifie si une potion existe dans l'inventaire
	index := -1
	for i, item := range c.Inventaire {
		if item == "Potion" {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("❌ Pas de potion dans l'inventaire !")
		return
	}

	// Retirer la potion de l'inventaire
	c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)

	// Récupérer 50 PV (sans dépasser PVMax)
	c.PVActuels += 50
	if c.PVActuels > c.PVMax {
		c.PVActuels = c.PVMax
	}

	fmt.Printf("✅ Vous avez utilisé une potion ! Points de vie : %d/%d\n", c.PVActuels, c.PVMax)
}
