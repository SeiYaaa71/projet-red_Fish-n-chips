package main

import "fmt"

func applyEffects(c *Character, m *Monster) {
	// Effets sur le joueur
	newEffets := []Effect{}
	for _, e := range c.Effets {
		switch e.Nom {
		case "Poison":
			c.PVActuels -= e.DamagePerTurn
			fmt.Printf(Red+"☠ Effet poison sur %s : -%d PV ! (PV : %d/%d)\n"+Reset, c.Nom, e.DamagePerTurn, c.PVActuels, c.PVMax)
		}
		e.Duration--
		if e.Duration > 0 {
			newEffets = append(newEffets, e)
		}
	}
	c.Effets = newEffets

	// Effets sur le monstre
	newEffetsM := []Effect{}
	for _, e := range m.Effets {
		switch e.Nom {
		case "Poison":
			m.PVActuels -= e.DamagePerTurn
			fmt.Printf(Red+"☠ Effet poison sur %s : -%d PV ! (PV : %d/%d)\n"+Reset, m.Nom, e.DamagePerTurn, m.PVActuels, m.PVMax)
		}
		e.Duration--
		if e.Duration > 0 {
			newEffetsM = append(newEffetsM, e)
		}
	}
	m.Effets = newEffetsM
}
