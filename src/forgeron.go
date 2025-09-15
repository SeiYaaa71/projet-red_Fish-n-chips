package main

import "fmt"

// Menu du Forgeron
func forgeron(c *Character) {
	for {
		clearScreen()
		fmt.Println(Bold + Blue + "\n=== FORGERON ===" + Reset)
		fmt.Printf("%sOr%s : %d pièces\n\n", Yellow, Reset, c.Gold)

		fmt.Println(Green + "1." + Reset + " Chapeau de l’aventurier (5 or)")
		fmt.Println(Green + "2." + Reset + " Tunique de l’aventurier (5 or)")
		fmt.Println(Green + "3." + Reset + " Bottes de l’aventurier (5 or)")
		fmt.Println(Red + "4." + Reset + " Retour au menu principal")
		fmt.Print(Yellow + "\nVotre choix : " + Reset)

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			craftItem(c, "Chapeau de l’aventurier", 5)
		case 2:
			craftItem(c, "Tunique de l’aventurier", 5)
		case 3:
			craftItem(c, "Bottes de l’aventurier", 5)
		case 4:
			return
		default:
			fmt.Println(Red + "❌ Choix invalide, réessayez." + Reset)
			waitForEnter()
		}
	}
}
