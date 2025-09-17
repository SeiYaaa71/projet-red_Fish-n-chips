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

	fmt.Println(Bold + Red + "\n⚔ Un gobelin apparaît pour l’entraînement !" + Reset)

	for !c.isDead() && !goblin.isDead() && !c.AbandonCombat {
		// --- Applique les effets (poison, etc.) au début du tour ---
		applyEffects(c, goblin)

		// Vérifie si quelqu’un est mort après les effets
		if c.isDead() || goblin.isDead() || c.AbandonCombat {
			break
		}

		// --- Tour du joueur ---
		characterTurn(c, goblin, &tour)

		if goblin.isDead() || c.AbandonCombat {
			break
		}

		// --- Tour du gobelin ---
		goblinPattern(goblin, c, tour)

		tour++
	}

	// --- Résultat du combat ---
	if goblin.isDead() {
		fmt.Println(Green + "\n🏆 Vous avez vaincu le Gobelin d’entraînement !" + Reset)

		// Gain d'expérience
		gainExp(c, goblin.ExpReward)

		// Récompense en or (aléatoire entre GoldMin et GoldMax)
		reward := rand.Intn(goblin.GoldMax-goblin.GoldMin+1) + goblin.GoldMin
		c.Gold += reward
		fmt.Printf(Yellow+"💰 Vous ramassez %d pièces d’or.\n"+Reset, reward)

	} else if c.isDead() {
		fmt.Println(Red + "\n💀 Vous avez été vaincu par le Gobelin..." + Reset)
	} else if c.AbandonCombat {
		fmt.Println(Red + "\n🏳 Vous avez abandonné le combat et perdu 10 or !" + Reset)
		if c.Gold >= 10 {
			c.Gold -= 10
		} else {
			c.Gold = 0
		}
	}
}
