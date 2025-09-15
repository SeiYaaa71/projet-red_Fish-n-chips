package main

import "fmt"

func menuEquipement(c *Character) {
	for {
		clearScreen()
		fmt.Println(Bold + Cyan + "\n=== MENU ÉQUIPEMENT ===" + Reset)
		displayEquipment(c)

		fmt.Println("\nOptions :")
		fmt.Println(Green + "1." + Reset + " Équiper un objet")
		fmt.Println(Red + "2." + Reset + " Retour au menu principal")

		var choix int
		fmt.Print(Yellow + "\nVotre choix : " + Reset)
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			if len(c.Inventaire) == 0 {
				fmt.Println(Red + "❌ Votre inventaire est vide !" + Reset)
				waitForEnter()
				continue
			}

			// Affiche seulement l’inventaire
			fmt.Println(Bold + Blue + "\n=== INVENTAIRE ===" + Reset)
			for i, item := range c.Inventaire {
				fmt.Printf(Green+"%d."+Reset+" %s\n", i+1, item)
			}

			fmt.Print(Yellow + "\nNuméro de l’objet à équiper : " + Reset)
			var index int
			fmt.Scanln(&index)

			if index < 1 || index > len(c.Inventaire) {
				fmt.Println(Red + "❌ Choix invalide." + Reset)
				waitForEnter()
				continue
			}

			item := c.Inventaire[index-1]
			equipItem(c, item)   // ⚡ équipe l’objet
			fmt.Println(Green + "✅ Objet équipé !" + Reset)
			waitForEnter()

		case 2:
			return // retour au menu principal

		default:
			fmt.Println(Red + "❌ Choix invalide." + Reset)
			waitForEnter()
		}
	}
}
