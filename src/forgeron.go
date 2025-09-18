package main

import "fmt"

// Menu du Forgeron
func forgeron(c *Character) {
	for {
		clearScreen()
		fmt.Println(Bold + Blue + "\n=== FORGERON ===" + Reset)
		fmt.Printf("%sOr%s : %d pièces\n\n", Yellow, Reset, c.Gold)

		// Équipements inspirés de Solo Leveling
		fmt.Println(Green + "1." + Reset + " Casque du Chasseur (5 or) [Nécessite 1x Plume de Corbeau]")
		fmt.Println(Green + "2." + Reset + " Armure du Chasseur (5 or) [Nécessite 2x Fourrure de Loup]")
		fmt.Println(Green + "3." + Reset + " Bottes du Chasseur (5 or) [Nécessite 1x Cuir de Sanglier]")
		fmt.Println(Green + "4." + Reset + " Forger une épée (+10 dmg) [Nécessite 2x Fer et 1x Bois + 10 or]")
		fmt.Println(Red + "5." + Reset + " Retour au menu principal")
		fmt.Print(Yellow + "\nVotre choix : " + Reset)

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			craftItem(c, "Casque du Chasseur", 5)
		case 2:
			craftItem(c, "Armure du Chasseur", 5)
		case 3:
			craftItem(c, "Bottes du Chasseur", 5)
		case 4:
			// Forge d’une épée (+10 dmg)
			if c.Gold >= 10 && hasItem(c, "Fer", 2) && hasItem(c, "Bois", 1) {
				c.Gold -= 10
				removeItem(c, "Fer", 2)
				removeItem(c, "Bois", 1)
				c.Inventaire = append(c.Inventaire, "Épée du Chasseur (+10 dmg)")
				c.Attaque += 10
				fmt.Println(Green + "✅ Vous avez forgé et équipé une Épée du Chasseur (+10 dmg) !" + Reset)
				waitForEnter()
			} else {
				fmt.Println(Red + "❌ Ressources insuffisantes (2x Fer, 1x Bois et 10 or requis)." + Reset)
				waitForEnter()
			}
		case 5:
			return
		default:
			fmt.Println(Red + "❌ Choix invalide, réessayez." + Reset)
			waitForEnter()
		}
	}
}
