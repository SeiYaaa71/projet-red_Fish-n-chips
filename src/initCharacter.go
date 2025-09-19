package main

func initCharacter(nom string, classe string) *Character {
	c := &Character{
		Nom:              nom,
		Classe:           classe,
		PVMax:            100, // ou selon la classe
		PVActuels:        50,
		Gold:             500,
		Inventaire:       []string{},
		Skill:            []string{"Coup de poing, Boule de feu, Soin mineur"},
		ExpActuelle:      0,
		ExpMax:           50,
		Niveau:           1,
		InventoryMax:     10,
		InventoryUpgrades: 0,
	}
	return c
}
