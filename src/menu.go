package main

import "fmt"

func menu(c *Character) {
	for {
		clearScreen()
		fmt.Println(Bold + Cyan + "=== Jeu RPG Console ===" + Reset)

		fmt.Println(Bold + Blue + "\n=== MENU PRINCIPAL ===" + Reset)
		fmt.Println(Green + "1." + Reset + " Afficher les informations du personnage")
		fmt.Println(Green + "2." + Reset + " Afficher l’inventaire")
		fmt.Println(Green + "3." + Reset + " Marchand")
		fmt.Println(Green + "4." + Reset + " Forgeron")
		fmt.Println(Green + "5." + Reset + " Équipement")
		fmt.Println(Green + "6." + Reset + " Combat d’entraînement")
		fmt.Println(Purple + "7." + Reset + " Mode Donjon")
		fmt.Println(Red + "8." + Reset + " Quitter le jeu")

		fmt.Print(Yellow + "\nVotre choix : " + Reset)

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			displayInfo(c)
			waitForEnter()
		case 2:
			accessInventory(c)
		case 3:
			merchant(c)
		case 4:
			forgeron(c)
		case 5:
			displayEquipment(c)
			waitForEnter()
		case 6:
			// Combat d’entraînement : gobelin classique
			trainingFight(c)
			waitForEnter()
		case 7:
			// Mode Donjon : système d’étages
			startDungeonRun(c)
		case 8:
			fmt.Println(Red + "👋 Merci d’avoir joué !" + Reset)
			return // Quitter le jeu
		default:
			fmt.Println(Red + "❌ Choix invalide, réessayez." + Reset)
			waitForEnter()
		}
	}
}
