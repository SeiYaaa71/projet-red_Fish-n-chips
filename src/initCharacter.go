package main

func initCharacter(nom string, classe string, niveau int, pvMax int, pvActuels int, inventaire []string) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      pvMax,
		PVActuels:  pvActuels,
		Inventaire: inventaire,
	}
}
