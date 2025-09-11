package main

import (
	"fmt"
)

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
		fmt.Println("3. Retour au menu principal")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			takePot(c)
		case 2:
			poisonPot(c)
		case 3:
			return
		default:
			fmt.Println("❌ Choix invalide, réessayez.")
		}
	}
}
