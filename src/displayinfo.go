package main

import "fmt"

// Fonction pour afficher les infos du personnage
func displayInfo(c Character) {
	fmt.Println("=== Informations du personnage ===")
	fmt.Println("Nom         :", c.Nom)
	fmt.Println("Classe      :", c.Classe)
	fmt.Println("Niveau      :", c.Niveau)
	fmt.Printf("Points de vie : %d/%d\n", c.PVActuels, c.PVMax)
	fmt.Println("Inventaire  :", c.Inventaire)
	fmt.Println("Compétences :", c.Skills)
	fmt.Println("Gold        :", c.Gold)
	fmt.Println("=================================")
}
