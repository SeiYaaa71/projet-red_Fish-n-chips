package main

import (
	"fmt"
	"os"
)

func menu(c *Character) {
	for {
		clearScreen()
		fmt.Println(Bold + Blue + "\n=== MENU PRINCIPAL ===" + Reset)
		fmt.Println(Green + "1." + Reset + " Afficher les informations du personnage")
		fmt.Println(Green + "2." + Reset + " Accéder au contenu de l’inventaire")
		fmt.Println(Green + "3." + Reset + " Marchand")
		fmt.Println(Green + "4." + Reset + " Forgeron")
		fmt.Println(Green + "5." + Reset + " Afficher l'équipement")
		fmt.Println(Green + "6." + Reset + " Combat d'entraînement")
		fmt.Println(Red + "7." + Reset + " Quitter")
		fmt.Print(Yellow + "\nVotre choix : " + Reset)

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			clearScreen()
			displayInfo(c)
			fmt.Println("\n(Appuyez sur Entrée pour revenir au menu)")
			waitForEnter()
		case 2:
			clearScreen()
			accessInventory(c)
		case 3:
			clearScreen()
			merchant(c)
		case 4:
			clearScreen()
			forgeron(c)
		case 5:
			clearScreen()
			displayEquipment(c)
			fmt.Println("\n(Appuyez sur Entrée pour revenir au menu)")
			waitForEnter()
		case 6: // Combat d'entraînement
			clearScreen()
			trainingFight(c) // fonction à créer pour gérer le combat
		case 7:
			fmt.Println(Red + "👋 Au revoir !" + Reset)
			os.Exit(0)
		default:
			fmt.Println(Red + "❌ Choix invalide, réessayez." + Reset)
			waitForEnter()
		}
	}
}
