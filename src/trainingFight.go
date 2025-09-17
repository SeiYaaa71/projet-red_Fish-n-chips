package main

import (
	"fmt"
	"math/rand"
	"time"
)

func trainingFight(c *Character) {
	goblin := initGoblin()
	turn := 1

	for goblin.PVActuels > 0 && c.PVActuels > 0 {
		clearScreen()
		fmt.Printf(Bold+"Tour %d\n"+Reset, turn)

		// Tour du joueur
		characterTurn(c, &goblin)

		// Vérif si le gobelin est mort
		if goblin.PVActuels <= 0 {
			fmt.Println(Green + "🎉 Vous avez vaincu le Gobelin d'entraînement !" + Reset)

			// Gain d'expérience
			gainExp(c, goblin.ExpReward)

			// Gain d’or aléatoire
			rand.Seed(time.Now().UnixNano())
			goldReward := rand.Intn(goblin.PVMax/2-goblin.PVMax/5+1) + goblin.PVMax/5
			c.Gold += goldReward
			fmt.Printf(Yellow+"💰 Vous avez gagné %d pièces d’or !\n"+Reset, goldReward)

			waitForEnter()
			return
		}

		// Tour du gobelin
		if c.PVActuels > 0 {
			if turn%3 == 0 {
				damage := goblin.Attaque * 2
				c.PVActuels -= damage
				fmt.Printf(Red+"%s inflige à %s %d dégâts (attaque puissante)!\n"+Reset,
					goblin.Nom, c.Nom, damage)
			} else {
				damage := goblin.Attaque
				c.PVActuels -= damage
				fmt.Printf(Red+"%s inflige à %s %d dégâts.\n"+Reset,
					goblin.Nom, c.Nom, damage)
			}

			if c.PVActuels < 0 {
				c.PVActuels = 0
			}

			fmt.Printf("❤️ %s : %d/%d PV\n", c.Nom, c.PVActuels, c.PVMax)
			waitForEnter()
		}

		turn++
	}
}
