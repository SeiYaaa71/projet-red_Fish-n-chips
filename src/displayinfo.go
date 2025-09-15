package main

import "fmt"

// Fonction pour afficher les infos du personnage
func displayInfo(c *Character) {
	fmt.Println("Nom :", c.Nom)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Niveau)
	fmt.Printf("PV : %d/%d\n", c.PVActuels, c.PVMax)
	fmt.Println("Inventaire :", c.Inventaire)
	fmt.Println("Sorts :", c.Skill)
	fmt.Println("Gold :", c.Gold)
}
