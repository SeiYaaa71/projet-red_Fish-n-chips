package main

import (
	"fmt"
	"os"
)

// Fonction pour afficher le menu principal
func menu(c *Character) {
	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l’inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Quitter")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			displayInfo(*c)

		case 2:
			accessInventory(c)

		case 3:
			merchant(c)

		case 4:
			fmt.Println("👋 Au revoir !")
			os.Exit(0)

		default:
			fmt.Println("❌ Choix invalide, réessayez.")
		}
	}
}
