package main

import "fmt"

// Vérifie si le joueur est mort
func isDead(c *Character) {
	if c.PVActuels <= 0 {
		fmt.Println("💀 Vous êtes mort...")

		// Résurrection avec 50% des PVMax
		c.PVActuels = c.PVMax / 2

		fmt.Printf("✨ Vous êtes ressuscité avec %d/%d PV.\n", c.PVActuels, c.PVMax)
	}
}
