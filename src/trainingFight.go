package main

import (
	"fmt"
	"math/rand"
)

// Combat d’entraînement contre un gobelin
func trainingFight(c *Character) {
	goblin := initGoblin() // *Monster
	tour := 1

	// Réinitialisation de l'abandon
	c.AbandonCombat = false

	// Liste des effets (poison, etc.)
	var monsterEffects []StatusEffect

	fmt.Println(Bold + Red + "\n⚔ Un gobelin apparaît pour l’entraînement !" + Reset)

	for !c.isDead() && !goblin.isDead() && !c.AbandonCombat {
		// --- Applique les effets (poison, etc.) au début du tour ---
		newEffects := []StatusEffect{}
		for _, eff := range monsterEffects {
			if eff.Type == "poison" {
				goblin.PVActuels -= eff.Damage
				fmt.Printf(Purple+"☠️ Le gobelin souffre du poison et perd %d PV ! (PV restants : %d/%d)\n"+Reset,
					eff.Damage, goblin.PVActuels, goblin.PVMax)
			}
			eff.Duration--
			if eff.Duration > 0 {
				newEffects = append(newEffects, eff)
			}
		}
		monsterEffects = newEffects

		// Vérifie si quelqu’un est mort après les effets
		if c.isDead() || goblin.isDead() || c.AbandonCombat {
			break
		}

		// --- Tour du joueur ---
		fmt.Printf(Bold+"\n=== Tour %d ===\n"+Reset, tour)
		fmt.Println(Green + "1." + Reset + " Attaquer")
		fmt.Println(Green + "2." + Reset + " Utiliser une compétence")
		fmt.Println(Green + "3." + Reset + " Utiliser une potion de vie")
		fmt.Println(Green + "4." + Reset + " Utiliser une potion de poison")
		fmt.Println(Red + "5." + Reset + " Abandonner le combat")
		fmt.Print(Yellow + "\nVotre choix : " + Reset)

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			// Attaque normale
			damage := rand.Intn(5) + 5
			goblin.PVActuels -= damage
			fmt.Printf(Green+"Vous attaquez le gobelin et infligez %d dégâts ! (PV restants : %d/%d)\n"+Reset,
				damage, goblin.PVActuels, goblin.PVMax)

		case 2:
			characterTurn(c, goblin, &tour)

		case 3:
			// Potion de vie
			useHealthPotion(c)

		case 4:
			// Potion de poison
			monsterEffects = append(monsterEffects, StatusEffect{
				Type:     "poison",
				Duration: 3,
				Damage:   10,
			})
			fmt.Println(Purple + "☠️ Vous lancez une potion de poison sur le gobelin !" + Reset)

		case 5:
			// Abandon
			fmt.Println(Red + "🏳 Vous abandonnez le combat !" + Reset)
			c.AbandonCombat = true
			if c.Gold >= 10 {
				c.Gold -= 10
			} else {
				c.Gold = 0
			}
			return

		default:
			fmt.Println(Red + "❌ Choix invalide !" + Reset)
			continue
		}

		// --- Tour du gobelin ---
		if !goblin.isDead() {
			goblinPattern(goblin, c, tour)
		}

		tour++
	}

	// --- Résultat du combat ---
	if goblin.isDead() {
		fmt.Println(Green + "\n🏆 Vous avez vaincu le Gobelin d’entraînement !" + Reset)

		// Gain d'expérience
		gainExp(c, goblin.ExpReward)

		// Récompense en or
		reward := rand.Intn(goblin.GoldMax-goblin.GoldMin+1) + goblin.GoldMin
		c.Gold += reward
		fmt.Printf(Yellow+"💰 Vous ramassez %d pièces d’or.\n"+Reset, reward)

	} else if c.isDead() {
		fmt.Println(Red + "\n💀 Vous avez été vaincu par le Gobelin..." + Reset)
	}
}
