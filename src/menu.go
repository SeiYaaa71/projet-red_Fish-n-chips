package main

import (
	"fmt"
	"os"
)

// Fonction pour afficher le menu principal
func menu(c *Character) {
	for {
		clearScreen() // Nettoie l’écran avant d’afficher le menu
		fmt.Println(Bold + Blue + "\n=== MENU PRINCIPAL ===" + Reset)
		fmt.Println(Green + "1." + Reset + " Afficher les informations du personnage")
		fmt.Println(Green + "2." + Reset + " Accéder au contenu de l’inventaire")
		fmt.Println(Green + "3." + Reset + " Marchand")
		fmt.Println(Red + "4." + Reset + " Quitter")
		fmt.Print(Yellow + "\nVotre choix : " + Reset)

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
			fmt.Println(Red + "👋 Au revoir !" + Reset)
			os.Exit(0)
		default:
			fmt.Println(Red + "❌ Choix invalide, réessayez." + Reset)
		}
	}
}
