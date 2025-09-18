package main

import "fmt"

func displayEquipment(c *Character) {
	fmt.Println("=== ÉQUIPEMENT ===")
	fmt.Println("🪖 Tête  :", equipOrNone(c.Equipement.Tete))
	fmt.Println("👕 Torse :", equipOrNone(c.Equipement.Torse))
	fmt.Println("🥾 Pieds :", equipOrNone(c.Equipement.Pieds))


	fmt.Printf("\n❤️ PV : %d / %d\n", c.PVActuels, c.PVMax)
	fmt.Printf("💥 ATK : %d\n", c.Attaque)
}

func equipOrNone(item string) string {
	if item == "" {
		return "Aucun"
	}
	return item
}
