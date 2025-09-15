package main

import "fmt"

// Fonction pour afficher l'équipement du personnage
func displayEquipment(c *Character) {
	fmt.Println("=== ÉQUIPEMENT ===")
	fmt.Printf("🪖 Tête  : %s\n", c.Equipement.Tete)
	fmt.Printf("👕 Torse : %s\n", c.Equipement.Torse)
	fmt.Printf("🥾 Pieds : %s\n", c.Equipement.Pieds)
}
