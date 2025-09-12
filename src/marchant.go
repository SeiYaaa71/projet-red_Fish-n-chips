package main

import "fmt"

// Interface du marchand
func merchant(c *Character) {
	for {
		fmt.Println("\n=== MARCHAND ===")
		fmt.Println("1. Acheter une Potion de vie (gratuit)")
		fmt.Println("2. Acheter une Potion de poison (gratuit)")
		fmt.Println("3. Acheter Livre de Sort : Boule de Feu (gratuit)")
		fmt.Println("4. Retour au menu principal")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			addInventory(c, "Potion")
		case 2:
			addInventory(c, "Potion de poison")
		case 3:
			addInventory(c, "Livre de Sort : Boule de Feu")
		case 4:
			return
		default:
			fmt.Println("❌ Choix invalide, réessayez.")
		}
	}
}
