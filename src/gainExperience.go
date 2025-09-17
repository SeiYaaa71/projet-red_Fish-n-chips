package main

import "fmt"

// Fonction qui ajoute de l'expérience au joueur et gère les niveaux
func gainExp(c *Character, expGagne int) {
	fmt.Printf(Green+"✨ Vous gagnez %d points d'expérience !\n"+Reset, expGagne)
	c.ExpActuelle += expGagne

	// Tant que l'expérience dépasse le max, on augmente le niveau
	for c.ExpActuelle >= c.ExpMax {
		c.ExpActuelle -= c.ExpMax
		c.Niveau++
		fmt.Printf(Yellow+"🎉 %s passe au niveau %d !\n"+Reset, c.Nom, c.Niveau)

		// Bonus de statistiques à chaque niveau
		c.PVMax += 10
		c.PVActuels = c.PVMax
		fmt.Printf(Green+"❤️ PV max augmenté à %d\n"+Reset, c.PVMax)

		// Augmentation progressive de l'expérience nécessaire pour le prochain niveau
		c.ExpMax = int(float64(c.ExpMax) * 1.5)
	}
	fmt.Printf(Cyan+"Expérience : %d/%d\n"+Reset, c.ExpActuelle, c.ExpMax)
	waitForEnter()
}
