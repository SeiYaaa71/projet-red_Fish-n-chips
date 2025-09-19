package main

import "fmt"

func displayInfo(c *Character) {
	fmt.Println("=== Informations du personnage ===")
	fmt.Printf("👤 Nom      : %s\n", c.Nom)
	fmt.Printf("⚔️  Classe   : %s\n", c.Classe)
	fmt.Printf("⭐ Niveau   : %d\n", c.Niveau)
	fmt.Printf("🧙‍♂️  Expérience : %d / %d\n", c.ExpActuelle, c.ExpMax)
	fmt.Printf("❤️  PV      : %d / %d\n", c.PVActuels, c.PVMax)
	fmt.Printf("💰 Or      : %d\n", c.Gold)
	fmt.Printf("🎒 Inventaire : %d / %d slots utilisés\n", len(c.Inventaire), c.InventoryMax)
	fmt.Printf("🔮 Compétences : %v\n", c.Skill)

	fmt.Println("\n=== Équipement ===")
	fmt.Printf("🪖  Tête  : %s\n", c.Equipement.Tete)
	fmt.Printf("👕 Torse : %s\n", c.Equipement.Torse)
	fmt.Printf("🥾 Pieds : %s\n", c.Equipement.Pieds)

	fmt.Printf("💥 Attaque : %d\n", c.Attaque)
}
