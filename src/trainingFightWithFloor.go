package main

import (
	"fmt"
	"math/rand"
)

// Combat dans le donjon avec étage
func trainingFightWithFloor(c *Character, m *Monster, floor int) {
	tour := 1
	c.AbandonCombat = false

	fmt.Printf(Bold+Red+"\n⚔ Un %s apparaît à l’étage %d !"+Reset+"\n", m.Nom, floor)

	for !c.isDead() && !m.isDead() {
		// Appliquer effets (poison, etc.)
		applyEffects(c, m)

		// Vérifie si quelqu’un est mort après les effets
		if c.isDead() || m.isDead() {
			break
		}

		// --- Tour du joueur ---
		fmt.Printf(Bold+"\n=== Tour %d ===\n"+Reset, tour)
		fmt.Println(Green + "1." + Reset + " Attaquer")
		fmt.Println(Green + "2." + Reset + " Utiliser une compétence")
		fmt.Println(Green + "3." + Reset + " Utiliser une potion de vie")
		fmt.Println(Red + "4." + Reset + " Abandonner le combat et Quitter le donjon")
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
		if len(c.Skill) == 0 {
		fmt.Println(Red + "❌ Vous n’avez pas de compétence à utiliser !" + Reset)
		} else {
		// Affiche la liste des compétences séparées
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
			// Utilisation d’une potion
			useHealthPotion(c)
		case 4:
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
