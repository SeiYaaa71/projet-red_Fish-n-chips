package main

import "fmt"

// Sous-menu inventaire
func accessInventory(c *Character) {
	for {
		clearScreen()
		fmt.Println(Bold + Blue + "\n=== INVENTAIRE ===" + Reset)

		if len(c.Inventaire) == 0 {
			fmt.Println(Red + "Inventaire vide" + Reset)
		} else {
			for i, item := range c.Inventaire {
				fmt.Printf(Green+"%d."+Reset+" %s\n", i+1, item)
			}
		}

		fmt.Println("\nOptions :")
		fmt.Println(Green + "1." + Reset + " Utiliser une Potion de vie")
		fmt.Println(Green + "2." + Reset + " Utiliser une Potion de poison")
		fmt.Println(Green + "3." + Reset + " Utiliser un Livre de Sort : Boule de Feu")
		fmt.Println(Red + "4." + Reset + " Retour au menu principal")
		fmt.Print(Yellow + "\nVotre choix : " + Reset)

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			takePot(c)
		case 2:
			poisonPot(c)
		case 3:
			spellBook(c)
		case 4:
			return
		default:
			fmt.Println(Red + "❌ Choix invalide, réessayez." + Reset)
		}
	}
}
