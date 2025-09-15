package main

import "fmt"

func trainingFight(c *Character) {
	gobelin := initGoblin()
	tour := 1

	for c.PVActuels > 0 && gobelin.PVActuels > 0 {
		fmt.Printf(Bold+Blue+"\n--- Tour %d ---\n"+Reset, tour)

		// Tour du joueur
		combatContinue := characterTurn(c, gobelin) // retourne false si abandon
		if !combatContinue {
			fmt.Println(Red + "⚔️ Vous quittez le combat !" + Reset)
			return // sortie immédiate du combat
		}

		// Si le gobelin est KO, on sort
		if gobelin.PVActuels <= 0 {
		fmt.Println(Green + "🎉 Vous avez vaincu le Gobelin d’entraînement !" + Reset)
		gainExperience(c, gobelin.ExpGagnee) // le joueur gagne de l'expérience
		return
		}

		// Tour du gobelin
		degats := gobelin.Attaque
		if tour%3 == 0 {
			degats *= 2
		}
		c.PVActuels -= degats
		if c.PVActuels < 0 {
			c.PVActuels = 0
		}

		fmt.Printf(Red+"%s inflige %d dégâts à %s !\n"+Reset,
			gobelin.Nom, degats, c.Nom)
		fmt.Printf("❤️ %s : %d/%d PV\n", c.Nom, c.PVActuels, c.PVMax)

		if c.PVActuels <= 0 {
			fmt.Println(Red + "💀 Vous êtes K.O. !" + Reset)
			return
		}

		tour++
		waitForEnter()
	}
}
