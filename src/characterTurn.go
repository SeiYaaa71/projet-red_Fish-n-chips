package main

import "fmt"

// Fonction qui simule le tour du joueur
func characterTurn(c *Character, m *Monster) bool { // retourne false si le joueur abandonne
	for {
		fmt.Println(Bold + Cyan + "\n=== Tour du joueur ===" + Reset)
		fmt.Println(Green + "1." + Reset + " Attaquer")
		fmt.Println(Green + "2." + Reset + " Inventaire")
		fmt.Println(Red + "3." + Reset + " Abandonner le combat")
		fmt.Print(Yellow + "\nVotre choix : " + Reset)

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1: // Attaque basique
			degats := 5
			m.PVActuels -= degats
			if m.PVActuels < 0 {
				m.PVActuels = 0
			}

			fmt.Printf(Green+"%s utilise Attaque basique et inflige %d dégâts à %s !\n"+Reset,
				c.Nom, degats, m.Nom)
			fmt.Printf("❤️ %s : %d/%d PV\n", m.Nom, m.PVActuels, m.PVMax)
			return true // le combat continue

		case 2: // Utilisation de l’inventaire
			if len(c.Inventaire) == 0 {
				fmt.Println(Red + "❌ Votre inventaire est vide !" + Reset)
				waitForEnter()
				continue
			}

			fmt.Println(Bold + "\n=== Inventaire ===" + Reset)
			for i, item := range c.Inventaire {
				fmt.Printf("%d. %s\n", i+1, item)
			}
			fmt.Println("0. Retour")

			fmt.Print("\nVotre choix : ")
			var choixItem int
			fmt.Scanln(&choixItem)

			if choixItem == 0 {
				continue
			}

			if choixItem < 1 || choixItem > len(c.Inventaire) {
				fmt.Println(Red + "❌ Choix invalide." + Reset)
				waitForEnter()
				continue
			}

			item := c.Inventaire[choixItem-1]

			switch item {
			case "Potion de vie":
				c.PVActuels += 50
				if c.PVActuels > c.PVMax {
					c.PVActuels = c.PVMax
				}
				fmt.Printf(Green+"✅ Vous utilisez %s ! PV : %d/%d\n"+Reset, item, c.PVActuels, c.PVMax)
			case "Potion de poison":
				fmt.Println(Red + "💀 Vous utilisez Potion de poison sur vous-même !" + Reset)
				poisonPot(c)
			case "Livre de Sort : Boule de Feu":
				if !contains(c.Skill, "Boule de Feu") {
					c.Skill = append(c.Skill, "Boule de Feu")
					fmt.Println(Green + "✨ Vous apprenez le sort Boule de Feu !" + Reset)
				} else {
					fmt.Println(Red + "❌ Vous connaissez déjà ce sort !" + Reset)
				}
			default:
				fmt.Printf(Red+"❌ L'objet %s n'est pas utilisable pour le moment.\n"+Reset, item)
			}

			// Supprimer l'objet utilisé de l'inventaire si consommable
			if item == "Potion de vie" || item == "Potion de poison" || item == "Livre de Sort : Boule de Feu" {
				removeInventory(c, item)
			}

			waitForEnter()
			return true

		case 3: // Abandonner
			perteOr := 10
			if c.Gold < perteOr {
				perteOr = c.Gold
			}
			c.Gold -= perteOr
			fmt.Printf(Red+"💨 Vous abandonnez le combat et perdez %d pièces d’or !\n"+Reset, perteOr)
			waitForEnter()
			return false // combat arrêté

		default:
			fmt.Println(Red + "❌ Choix invalide, réessayez." + Reset)
		}
	}
}

// Fonction utilitaire pour vérifier si une slice contient un élément
func contains(slice []string, elem string) bool {
	for _, s := range slice {
		if s == elem {
			return true
		}
	}
	return false
}
