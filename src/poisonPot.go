package main

import (
	"fmt"
	"time"
)

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
