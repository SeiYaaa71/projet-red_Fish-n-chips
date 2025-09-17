package main

import (
	"fmt"
	"time"
)

// Utilise une potion de soin si disponible
func useHealthPotion(c *Character) {
	index := -1
	for i, item := range c.Inventaire {
		if item == "Potion de vie" {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println(Red + "❌ Vous n'avez pas de potion de vie dans votre inventaire !" + Reset)
		return
	}

	// Retire la potion de l'inventaire
	c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)

	// Restaure les PV du personnage
	heal := 20
	c.PVActuels += heal
	if c.PVActuels > c.PVMax {
		c.PVActuels = c.PVMax
	}

	fmt.Printf(Green+"💚 Vous utilisez une Potion de vie et récupérez %d PV ! (PV : %d/%d)\n"+Reset,
		heal, c.PVActuels, c.PVMax)
}

func poisonPot(c *Character) {
	fmt.Println("☠️ Vous avez bu une Potion de poison !")

	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)

		c.PVActuels -= 10
		if c.PVActuels < 0 {
			c.PVActuels = 0
		}
		fmt.Printf("⏳ Seconde %d : %d/%d PV\n", i, c.PVActuels, c.PVMax)
		isDead(c)
	}
}
