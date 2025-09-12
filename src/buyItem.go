package main

import (
	"fmt"
	"bufio"
	"os"
)

// Fonction générique pour acheter un item
func buyItem(c *Character, item string, price int) {
	if c.Gold < price {
		fmt.Printf(Red+"❌ Vous n'avez pas assez d'or pour acheter %s.\nIl vous faut %d or, mais vous n'avez que %d or.\n"+Reset,
			item, price, c.Gold)
		waitForEnter()
		return
	}

	if canAddItem(c) {
		c.Gold -= price
		c.Inventaire = append(c.Inventaire, item)
		fmt.Printf(Green+"✅ Vous avez acheté %s pour %d or.\n"+Reset, item, price)
		waitForEnter()
	}
}

// Pause jusqu'à ce que le joueur appuie sur Entrée
func waitForEnter() {
	fmt.Print("\n" + Cyan + "Appuyez sur Entrée pour continuer..." + Reset)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
