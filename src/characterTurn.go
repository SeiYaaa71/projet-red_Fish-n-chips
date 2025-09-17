package main

import (
	"fmt"
)

// characterTurn simule le tour du joueur
func characterTurn(c *Character, m *Monster, tour *int) {
	for {
		fmt.Println(Bold + Cyan + "\n=== VOTRE TOUR ===" + Reset)
		fmt.Println(Green + "1."+ Reset + "Attaquer")
		fmt.Println(Green + "2."+ Reset + "Utiliser un sort")
		fmt.Println(Green + "3."+ Reset + "Utiliser un objet")
		fmt.Println(Green + "4."+ Reset + "Abandonner le combat")
		fmt.Print(Yellow + "\nVotre choix : " + Reset)

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1: // Attaque basique
			damage := 5
			m.PVActuels -= damage
			if m.PVActuels < 0 {
				m.PVActuels = 0
			}
			fmt.Printf(Green+"%s inflige %d dégâts à %s ! (PV %d/%d)\n"+Reset,
				c.Nom, damage, m.Nom, m.PVActuels, m.PVMax)
			return // Fin du tour du joueur

		case 2: // Sorts
			fmt.Println("\n=== SORTS ===")
			fmt.Println("1. Coup de poing (8 dégâts)")
			fmt.Println("2. Boule de feu (18 dégâts)")
			fmt.Println("3. Sort de soin (+20 PV)")
			fmt.Print(Yellow + "\nVotre choix : " + Reset)

			var sortChoix int
			fmt.Scanln(&sortChoix)

			switch sortChoix {
			case 1:
				damage := 8
				m.PVActuels -= damage
				if m.PVActuels < 0 {
					m.PVActuels = 0
				}
				fmt.Printf(Green+"%s utilise Coup de poing et inflige %d dégâts à %s ! (PV %d/%d)\n"+Reset,
					c.Nom, damage, m.Nom, m.PVActuels, m.PVMax)
			case 2:
				damage := 18
				m.PVActuels -= damage
				if m.PVActuels < 0 {
					m.PVActuels = 0
				}
				fmt.Printf(Green+"%s lance Boule de feu et inflige %d dégâts à %s ! (PV %d/%d)\n"+Reset,
					c.Nom, damage, m.Nom, m.PVActuels, m.PVMax)
			case 3:
				heal := 20
				c.PVActuels += heal
				if c.PVActuels > c.PVMax {
					c.PVActuels = c.PVMax
				}
				fmt.Printf(Green+"%s utilise un sort de soin et récupère %d PV ! (PV %d/%d)\n"+Reset,
					c.Nom, heal, c.PVActuels, c.PVMax)
			default:
				fmt.Println(Red + "❌ Sort invalide !" + Reset)
				continue
			}
			return // Fin du tour après utilisation d'un sort

		case 3: // Objets
			if len(c.Inventaire) == 0 {
				fmt.Println(Red + "❌ Votre inventaire est vide !" + Reset)
				return
			}
			fmt.Println("\n=== OBJETS ===")
			for i, item := range c.Inventaire {
				fmt.Printf("%d. %s\n", i+1, item)
			}
			fmt.Print(Yellow + "\nNuméro de l’objet à utiliser : " + Reset)

			var index int
			fmt.Scanln(&index)
			if index < 1 || index > len(c.Inventaire) {
				fmt.Println(Red + "❌ Choix invalide." + Reset)
				continue
			}

			item := c.Inventaire[index-1]
			switch item {
			case "Potion de vie":
				useHealthPotion(c)
			case "Potion de poison":
				fmt.Println("1. Se jeter la potion de poison")
				fmt.Println("2. Lancer la potion de poison sur le monstre")
				fmt.Print(Yellow + "\nVotre choix : " + Reset)
				var pChoix int
				fmt.Scanln(&pChoix)
				if pChoix == 1 {
					usePoison(c)
				} else if pChoix == 2 {
					throwPoisonOnMonster(m)
				} else {
					fmt.Println(Red + "❌ Choix invalide !" + Reset)
					continue
				}
			default:
				fmt.Println(Red + "❌ Objet non utilisable en combat !" + Reset)
				continue
			}

			// Retirer l'objet de l'inventaire après usage
			c.Inventaire = append(c.Inventaire[:index-1], c.Inventaire[index:]...)
			return // Fin du tour après utilisation d’un objet

		case 4:
			// Abandon
			fmt.Println(Red + "\n🏳 Vous abandonnez le combat !" + Reset)
			c.Gold -= 10
			if c.Gold < 0 {
				c.Gold = 0
			}
			// On peut gérer un flag pour sortir de la boucle combat
			c.AbandonCombat = true
			return

		default:
			fmt.Println(Red + "❌ Choix invalide !" + Reset)
		}
	}
}
