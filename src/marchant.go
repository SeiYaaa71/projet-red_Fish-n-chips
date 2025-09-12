package main

import "fmt"

// Interface du marchand avec système de gold
func merchant(c *Character) {
	for {
		clearScreen()
		fmt.Println(Bold + Blue + "\n=== MARCHAND ===" + Reset)
		fmt.Printf("%sOr%s : %d pièces\n\n", Yellow, Reset, c.Gold)
		
		fmt.Println(Green + "1." + Reset + " Acheter une Potion de vie (20 or)")
		fmt.Println(Green + "2." + Reset + " Acheter une Potion de poison (15 or)")
		fmt.Println(Green + "3." + Reset + " Acheter un Livre de Sort : Boule de Feu (50 or)")
		fmt.Println(Red + "4." + Reset + " Retour au menu principal")
		fmt.Print(Yellow + "\nVotre choix : " + Reset)

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			buyItem(c, "Potion", 20)
		case 2:
			buyItem(c, "Potion de poison", 15)
		case 3:
			buyItem(c, "Livre de Sort : Boule de Feu", 50)
		case 4:
			return
		default:
			fmt.Println(Red + "❌ Choix invalide, réessayez." + Reset)
		}
	}
}
