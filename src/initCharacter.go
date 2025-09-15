package main

// Fonction pour initialiser un personnage
func initCharacter(
	nom string,
	classe string,
	niveau int,
	pvMax int,
	pvActuels int,
	inventaire []string,
	skill []string,
	gold int,
) *Character {
	return &Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      pvMax,
		PVActuels:  pvActuels,
		Inventaire: inventaire,
		Skill:     skill,
		Gold:      gold,
		Equipement: Equipment{
			Tete:  "Aucun",
			Torse: "Aucun",
			Pieds: "Aucun",
		},
	}
}
