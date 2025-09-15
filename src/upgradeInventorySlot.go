package main

import "fmt"

func upgradeInventorySlot(c *Character) {
	if c.InventoryUpgrades >= 3 {
		fmt.Println(Red + "❌ Vous ne pouvez plus améliorer votre inventaire !" + Reset)
		waitForEnter()
		return
	}

	c.InventoryMax += 10
	c.InventoryUpgrades++
	fmt.Printf(Green+"✅ Inventaire augmenté ! Nouvelle capacité : %d\n"+Reset, c.InventoryMax)
	waitForEnter()
}
