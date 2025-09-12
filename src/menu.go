package main

import (
	"fmt"
	"os"
)

// Fonction pour afficher le menu principal
func menu(c *Character) {
	for {
		clearScreen() // Nettoie l’écran avant d’afficher le menu
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
			clearScreen()
			displayInfo(*c)
			fmt.Println("\n(Appuyez sur Entrée pour revenir au menu)")
			fmt.Scanln()
		case 2:
			clearScreen()
			accessInventory(c)
		case 3:
			clearScreen()
			merchant(c)
		case 4:
			fmt.Println("👋 Au revoir !")
			os.Exit(0)
		default:
			fmt.Println("❌ Choix invalide, réessayez.")
		}
	}
}
