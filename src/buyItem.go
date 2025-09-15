package main

import (
	"fmt"
)

// Fonction générique pour acheter un item
func buyItem(c *Character, item string, price int) {
	if c.Gold < price {
		fmt.Printf(Red+"❌ Vous n'avez pas assez d'or pour acheter %s.\nIl vous faut %d or, mais vous n'avez que %d or.\n"+Reset,
			item, price, c.Gold)
		waitForEnter()
		return
	}

	if len(c.Inventaire) >= c.InventoryMax {
		fmt.Println(Red + "❌ Votre inventaire est plein, impossible d’acheter cet objet." + Reset)
		waitForEnter()
		return
	}

	c.Gold -= price
	c.Inventaire = append(c.Inventaire, item)
	fmt.Printf(Green+"✅ Vous avez acheté %s pour %d or.\n"+Reset, item, price)
	waitForEnter()
}
