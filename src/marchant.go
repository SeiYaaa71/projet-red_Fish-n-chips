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
		fmt.Println(Green + "3." + Reset + " Fourrure de Loup (4 or)")
		fmt.Println(Green + "4." + Reset + " Peau de Goblin (7 or)")
		fmt.Println(Green + "5." + Reset + " Cuir de Sanglier (3 or)")
		fmt.Println(Green + "6." + Reset + " Plume de Corbeau (1 or)")
		fmt.Println(Green + "7." + Reset + " Fer (30 or)")
		fmt.Println(Green + "8." + Reset + " Bois (20 or)")
		fmt.Println(Green + "10." + Reset + " Vendre un objet")

		// ✅ Affiche l’augmentation d’inventaire seulement si < 3 upgrades
		if c.InventoryUpgrades < 3 {
			fmt.Printf(Green+"9."+Reset+" Augmenter la taille de l’inventaire (30 or) [%d/3]\n", c.InventoryUpgrades)
			fmt.Println(Red + "0." + Reset + " Retour au menu principal")
		} else {
			fmt.Println(Red + "0." + Reset + " Retour au menu principal")
		}

		fmt.Print(Yellow + "\nVotre choix : " + Reset)

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			buyItem(c, "Potion de vie", 3)
		case 2:
			buyItem(c, "Potion de poison", 6)
		case 3:
			buyItem(c, "Fourrure de Loup", 4)
		case 4:
			buyItem(c, "Peau de Goblin", 7)
		case 5:
			buyItem(c, "Cuir de Sanglier", 3)
		case 6:
			buyItem(c, "Plume de Corbeau", 1)
		case 7:
			if c.Gold >= 30 {
				c.Gold -= 30
				c.Inventaire = append(c.Inventaire, "Fer")
				fmt.Println("✅ Vous avez acheté du Fer.")
			} else {
				fmt.Println("❌ Pas assez d’or.")
			}
		case 8:
			if c.Gold >= 20 {
				c.Gold -= 20
				c.Inventaire = append(c.Inventaire, "Bois")
				fmt.Println("✅ Vous avez acheté du Bois.")
			} else {
				fmt.Println("❌ Pas assez d’or.")
				waitForEnter()
			}
		case 9:
			if c.InventoryUpgrades < 3 {
				if c.Gold >= 30 {
					c.Gold -= 30
					upgradeInventorySlot(c)
					fmt.Printf(Green+"✅ Inventaire amélioré : %d/3\n"+Reset, c.InventoryUpgrades)
					waitForEnter()
				} else {
					fmt.Println(Red + "❌ Pas assez d'or pour augmenter l’inventaire." + Reset)
					waitForEnter()
				}
			} else {
				// si déjà max, le 9 devient retour
				return
			}
			case 10:
			sellItem(c)
		case 0:
			return
		default:
			fmt.Println(Red + "❌ Choix invalide, réessayez." + Reset)
			waitForEnter()
		}
	}
}
