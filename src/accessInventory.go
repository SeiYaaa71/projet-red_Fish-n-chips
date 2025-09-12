package main

import (
	"fmt"
)

// Sous-menu inventaire
func accessInventory(c *Character) {
	for {
		fmt.Println("\n=== INVENTAIRE ===")
		if len(c.Inventaire) == 0 {
			fmt.Println("Inventaire vide")
		} else {
			for i, item := range c.Inventaire {
				fmt.Printf("%d. %s\n", i+1, item)
			}
		}

		fmt.Println("\nOptions :")
		fmt.Println("1. Utiliser une Potion de vie")
		fmt.Println("2. Utiliser une Potion de poison")
		fmt.Println("3. Utiliser Livre de Sort : Boule de Feu")
		fmt.Println("4. Retour au menu principal")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			takePot(c)
		case 2:
			poisonPot(c)
		case 3:
			// Vérifie si le joueur a le livre dans son inventaire
			index := -1
			for i, item := range c.Inventaire {
				if item == "Livre de Sort : Boule de Feu" {
					index = i
					break
				}
			}
			if index == -1 {
				fmt.Println("❌ Vous n’avez pas ce livre dans votre inventaire.")
			} else {
				spellBook(c)
				removeInventory(c, index) // Retire le livre après usage
			}
		case 4:
			return
		default:
			fmt.Println("❌ Choix invalide, réessayez.")
		}
	}
}
