package main

import "fmt"

func takePot(c *Character, potionName string) {
	// Vérifie d'abord si l'inventaire est vide
	if len(c.Inventaire) == 0 {
		fmt.Println(Red + "❌ Votre inventaire est vide !" + Reset)
		waitForEnter()
		return
	}

	// Cherche la potion dans l'inventaire
	found := false
	for _, item := range c.Inventaire {
		if item == potionName {
			found = true
			break
		}
	}

	if !found {
		fmt.Printf(Red+"❌ Vous n'avez pas de %s dans votre inventaire !\n"+Reset, potionName)
		waitForEnter()
		return
	}

	// Applique l'effet de la potion
	switch potionName {
	case "Potion de vie":
		c.PVActuels += 50
		if c.PVActuels > c.PVMax {
			c.PVActuels = c.PVMax
		}
		fmt.Printf(Green+"✅ Vous avez utilisé une Potion de vie ! PV : %d/%d\n"+Reset, c.PVActuels, c.PVMax)
	case "Potion de poison":
		fmt.Println(Red + "💀 Vous avez bu une potion de poison !" + Reset)
		poisonPot(c)
	default:
		fmt.Printf(Red+"❌ Potion inconnue : %s\n"+Reset, potionName)
		waitForEnter()
		return
	}

	// Supprime la potion de l'inventaire
	removeInventory(c, potionName)

	waitForEnter()
}
