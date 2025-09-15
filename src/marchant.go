package main

import "fmt"

// Interface du marchand avec système de gold
func merchant(c *Character) {
	for {
		clearScreen()
		fmt.Println(Bold + Blue + "\n=== MARCHAND ===" + Reset)
		fmt.Printf("%sOr%s : %d pièces\n\n", Yellow, Reset, c.Gold)

		// Liste des items en vente
		fmt.Println(Green + "1." + Reset + " Potion de vie (3 or)")
		fmt.Println(Green + "2." + Reset + " Potion de poison (6 or)")
		fmt.Println(Green + "3." + Reset + " Livre de Sort : Boule de Feu (25 or)")
		fmt.Println(Green + "4." + Reset + " Fourrure de Loup (4 or)")
		fmt.Println(Green + "5." + Reset + " Peau de Troll (7 or)")
		fmt.Println(Green + "6." + Reset + " Cuir de Sanglier (3 or)")
		fmt.Println(Green + "7." + Reset + " Plume de Corbeau (1 or)")
		fmt.Println(Red + "8." + Reset + " Retour au menu principal")
		fmt.Print(Yellow + "\nVotre choix : " + Reset)

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			buyItem(c, "Potion de vie", 3)
		case 2:
			buyItem(c, "Potion de poison", 6)
		case 3:
			buyItem(c, "Livre de Sort : Boule de Feu", 25)
		case 4:
			buyItem(c, "Fourrure de Loup", 4)
		case 5:
			buyItem(c, "Peau de Troll", 7)
		case 6:
			buyItem(c, "Cuir de Sanglier", 3)
		case 7:
			buyItem(c, "Plume de Corbeau", 1)
		case 8:
			return
		default:
			fmt.Println(Red + "❌ Choix invalide, réessayez." + Reset)
			waitForEnter()
		}
	}
}
