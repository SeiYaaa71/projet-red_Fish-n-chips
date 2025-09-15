package main

import "fmt"

func displayEquipment(c *Character) {
	fmt.Println("=== ÉQUIPEMENT ===")
	fmt.Println("🪖 Tête  :", equipOrNone(c.Equipement.Tete))
	fmt.Println("👕 Torse :", equipOrNone(c.Equipement.Torse))
	fmt.Println("🥾 Pieds :", equipOrNone(c.Equipement.Pieds))
	fmt.Printf("\n❤️ PV : %d / %d\n", c.PVActuels, c.PVMax)
}

func equipOrNone(slot string) string {
	if slot == "" {
		return "Aucun"
	}
	return slot
}