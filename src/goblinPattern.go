package main

import "fmt"

// Pattern d’attaque du gobelin
func goblinPattern(m *Monster, c *Character, turn int) {
	damage := m.Attaque

	// Tous les 3 tours → 200% des dégâts
	if turn%3 == 0 {
		damage *= 2
	}

	c.PVActuels -= damage
	if c.PVActuels < 0 {
		c.PVActuels = 0
	}

	fmt.Printf(Red+"%s inflige %d dégâts à %s !\n"+Reset, m.Nom, damage, c.Nom)
	fmt.Printf("❤️ %s : %d/%d PV\n\n", c.Nom, c.PVActuels, c.PVMax)
	waitForEnter()
}
