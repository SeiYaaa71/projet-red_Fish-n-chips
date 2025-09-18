package main

import (
	"fmt"
	"math/rand"
)

// Structure pour gérer les effets temporaires
type StatusEffect struct {
	Type     string // "poison"
	Duration int
	Damage   int
}

func trainingFightWithFloor(c *Character, m *Monster, floor int) {
	tour := 1
	c.AbandonCombat = false

	// Liste des effets actifs sur le monstre
	var monsterEffects []StatusEffect

	fmt.Printf(Bold+Red+"\n⚔ Un %s apparaît à l’étage %d !"+Reset+"\n", m.Nom, floor)

	for !c.isDead() && !m.isDead() {
		// --- Appliquer effets (poison, etc.) ---
		newEffects := []StatusEffect{}
		for _, eff := range monsterEffects {
			if eff.Type == "poison" {
				m.PVActuels -= eff.Damage
				fmt.Printf(Purple+"☠️ %s souffre du poison et perd %d PV ! (PV restants : %d/%d)\n"+Reset,
					m.Nom, eff.Damage, m.PVActuels, m.PVMax)
			}
			eff.Duration--
			if eff.Duration > 0 {
				newEffects = append(newEffects, eff)
			}
		}
		monsterEffects = newEffects

		// Vérifie si le monstre est mort après les effets
		if m.isDead() {
			break
		}

		// --- Tour du joueur ---
		fmt.Printf(Bold+"\n=== Tour %d ===\n"+Reset, tour)
		fmt.Println(Green + "1." + Reset + " Attaquer")
		fmt.Println(Green + "2." + Reset + " Utiliser une compétence")
		fmt.Println(Green + "3." + Reset + " Utiliser une potion de vie")
		fmt.Println(Green + "4." + Reset + " Utiliser une potion de poison")
		fmt.Println(Red + "5." + Reset + " Abandonner le combat et Quitter le donjon")
		fmt.Print(Yellow + "\nVotre choix : " + Reset)

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			// Attaque normale
			damage := rand.Intn(5) + 5
			m.PVActuels -= damage
			fmt.Printf(Green+"Vous attaquez %s et infligez %d dégâts ! (PV restants : %d/%d)\n"+Reset,
				m.Nom, damage, m.PVActuels, m.PVMax)

		case 2:
			// Utiliser une compétence
			if len(c.Skill) == 0 {
				fmt.Println(Red + "❌ Vous n’avez pas de compétence à utiliser !" + Reset)
			} else {
				fmt.Println(Green + "\nCompétences disponibles :" + Reset)
				fmt.Println("1. Coup de poing (8 Dmg)")
				fmt.Println("2. Boule de Feu (18 Dmg)")
				fmt.Println("3. Soin Mineur (15 PV)")
				fmt.Print(Yellow + "\nChoisissez une compétence : " + Reset)

				var skillChoice int
				fmt.Scanln(&skillChoice)

				switch skillChoice {
				case 1: // Coup de poing
					damage := 8
					m.PVActuels -= damage
					fmt.Printf(Green+"Vous utilisez Coup de poing et infligez %d dégâts à %s ! (PV restants : %d/%d)\n"+Reset,
						damage, m.Nom, m.PVActuels, m.PVMax)
				case 2: // Boule de feu
					damage := 18
					m.PVActuels -= damage
					fmt.Printf(Green+"Vous lancez Boule de Feu et infligez %d dégâts à %s ! (PV restants : %d/%d)\n"+Reset,
						damage, m.Nom, m.PVActuels, m.PVMax)
				case 3: // Soin mineur
					heal := 15
					c.PVActuels += heal
					if c.PVActuels > c.PVMax {
						c.PVActuels = c.PVMax
					}
					fmt.Printf(Green+"Vous utilisez Soin Mineur et récupérez %d PV ! (PV : %d/%d)\n"+Reset,
						heal, c.PVActuels, c.PVMax)
				default:
					fmt.Println(Red + "❌ Choix invalide !" + Reset)
					continue
				}
			}

		case 3:
			// Utilisation potion de vie
			useHealthPotion(c)

		case 4:
			// Potion de poison : applique effet pendant 3 tours
			monsterEffects = append(monsterEffects, StatusEffect{
				Type:     "poison",
				Duration: 3,
				Damage:   10,
			})
			fmt.Printf(Purple+"☠️ Vous lancez une potion de poison sur %s ! Il prendra %d dégâts par tour pendant %d tours.\n"+Reset,
				m.Nom, 10, 3)

		case 5:
			// Abandon
			fmt.Println(Red + "❌ Vous abandonnez le combat !" + Reset)
			c.AbandonCombat = true
			return

		default:
			fmt.Println(Red + "❌ Choix invalide, réessayez." + Reset)
			continue
		}

		// --- Tour du monstre ---
		if !m.isDead() {
			damage := rand.Intn(m.Attaque/2) + m.Attaque/2
			c.PVActuels -= damage
			fmt.Printf(Red+"%s attaque et vous inflige %d dégâts ! (PV : %d/%d)\n"+Reset,
				m.Nom, damage, c.PVActuels, c.PVMax)
		}

		tour++
	}

	// --- Résultat ---
	if m.isDead() {
		fmt.Println(Green + "\n🏆 Vous avez vaincu le monstre !" + Reset)

		// Récompense
		expGain := m.ExpReward
		goldGain := rand.Intn(m.GoldMax-m.GoldMin+1) + m.GoldMin
		c.ExpActuelle += expGain
		c.Gold += goldGain
		fmt.Printf(Yellow+"💰 Vous gagnez %d or et %d EXP !\n"+Reset, goldGain, expGain)
	} else if c.isDead() {
		fmt.Println(Red + "\n💀 Vous êtes mort..." + Reset)
	}
}
