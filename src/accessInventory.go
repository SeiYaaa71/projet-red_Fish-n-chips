package main

import "fmt"

func accessInventory(c *Character) {
	for {
		clearScreen()
		fmt.Println(Bold + Blue + "\n=== INVENTAIRE ===" + Reset)
		if len(c.Inventaire) == 0 {
			fmt.Println(Red + "Inventaire vide" + Reset)
		} else {
			fmt.Println(Green + "Vos objets :" + Reset)
			for i, item := range c.Inventaire {
				fmt.Printf("%d. %s\n", i+1, item)
			}
		}

		fmt.Println("\nOptions :")
		fmt.Println(Green + "1." + Reset + " Utiliser une Potion de vie")
		fmt.Println(Green + "2." + Reset + " Utiliser une Potion de poison")
		fmt.Println(Red + "3." + Reset + " Retour au menu principal")
		fmt.Print(Yellow + "\nVotre choix : " + Reset)

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			takePot(c, "Potion de vie")
		case 2:
			takePot(c, "Potion de poison")
		case 3:
			return
		default:
			fmt.Println(Red + "❌ Choix invalide, réessayez." + Reset)
			waitForEnter()
		}
	}
}
