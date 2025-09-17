package main

func initGoblin() *Monster {
	return &Monster{
		Nom:       "Gobelin d’entraînement",
		PVMax:     40,
		PVActuels: 40,
		Attaque:   5,
		ExpReward: 20,
	}
}
