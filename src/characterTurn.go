package main

import (
	"fmt"
	"strings"
)

// Tour du joueur
func characterTurn(c *Character, m *Monster, tour *int) {
	for {
		clearScreen()
		fmt.Printf(Bold+Blue+"--- Tour %d ---\n"+Reset, *tour)
		fmt.Printf("❤️ PV Joueur : %d/%d | 💀 PV Ennemi : %d/%d\n", c.PVActuels, c.PVMax, m.PVActuels, m.PVMax)
		fmt.Println("\nActions disponibles :")
		fmt.Println(Green + "1." + Reset + " Attaquer (coup classique)")
		fmt.Println(Green + "2." + Reset + " Utiliser un sort")
		fmt.Println(Green + "3." + Reset + " Utiliser un objet")
		fmt.Println(Red + "4." + Reset + " Fuir le combat")

		var choix int
		fmt.Print(Yellow + "\nVotre choix : " + Reset)
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			// Attaque classique
			damage := 10
			fmt.Printf("🗡 Vous attaquez et infligez %d dégâts au %s !\n", damage, m.Nom)
			m.PVActuels -= damage
			waitForEnter()
			return

		case 2:
			// Menu des sorts
			fmt.Println("\n--- Sorts disponibles ---")
			fmt.Println(Green + "1." + Reset + " Coup de poing (8 dégâts)")
			fmt.Println(Green + "2." + Reset + " Boule de feu (18 dégâts)")
			fmt.Println(Green + "3." + Reset + " Retour")

			var spell int
			fmt.Print(Yellow + "\nVotre choix : " + Reset)
			fmt.Scanln(&spell)

			switch spell {
			case 1:
				fmt.Println("👊 Vous utilisez Coup de poing et infligez 8 dégâts !")
				m.PVActuels -= 8
			case 2:
				fmt.Println("🔥 Vous lancez une Boule de feu et infligez 18 dégâts !")
				m.PVActuels -= 18
			case 3:
				continue
			default:
				fmt.Println(Red + "❌ Choix invalide." + Reset)
			}
			waitForEnter()
			return

		case 3:
			// Inventaire
			if len(c.Inventaire) == 0 {
				fmt.Println(Red + "❌ Votre inventaire est vide." + Reset)
				waitForEnter()
				continue
			}

			fmt.Println("\n--- Inventaire ---")
			for i, item := range c.Inventaire {
				fmt.Printf("%d. %s\n", i+1, item)
			}
			fmt.Printf("%d. Retour\n", len(c.Inventaire)+1)

			var choixItem int
			fmt.Print(Yellow + "\nVotre choix : " + Reset)
			fmt.Scanln(&choixItem)

			if choixItem < 1 || choixItem > len(c.Inventaire)+1 {
				fmt.Println(Red + "❌ Choix invalide." + Reset)
				waitForEnter()
				continue
			}
			if choixItem == len(c.Inventaire)+1 {
				continue
			}

			item := c.Inventaire[choixItem-1]

			// Utilisation des objets
			if strings.Contains(item, "Potion de vie") {
				c.PVActuels += 20
				if c.PVActuels > c.PVMax {
					c.PVActuels = c.PVMax
				}
				fmt.Println(Green + "🍷 Vous buvez une potion de vie et regagnez 20 PV." + Reset)

			} else if strings.Contains(item, "Potion de poison") {
				fmt.Printf(Green+"☠ Vous jetez une potion de poison sur %s !\n"+Reset, m.Nom)
				m.Effets = append(m.Effets, Effect{Nom: "Poison", Duration: 3, DamagePerTurn: 5})

			} else {
				fmt.Println(Red + "❌ Cet objet n’a aucun effet utilisable en combat." + Reset)
				waitForEnter()
				continue
			}

			// Retire l’objet utilisé
			c.Inventaire = append(c.Inventaire[:choixItem-1], c.Inventaire[choixItem:]...)
			waitForEnter()
			return

		case 4:
			// Fuite
			fmt.Println(Red + "🏳 Vous tentez de fuir le combat !" + Reset)
			c.PVActuels = 0 // considéré comme abandon
			waitForEnter()
			return

		default:
			fmt.Println(Red + "❌ Choix invalide." + Reset)
			waitForEnter()
		}
	}
}
