package main

import (
	"math/rand"
)
// Récupère les données d’un étage donné
func getFloorData(floor int) DungeonFloor {
	return DungeonFloor{
		Numero:   floor,
		ExpMult:  1.0 + float64(floor-1)*0.2, // +20% exp par étage
		GoldMult: 1.0 + float64(floor-1)*0.2, // +20% gold par étage
		HPMult:   1.0 + float64(floor-1)*0.3, // +30% PV par étage
		DmgMult:  1.0 + float64(floor-1)*0.2, // +20% dégâts par étage
	}
}

// Génère un monstre en fonction de l’étage
func initMonsterForFloor(floor DungeonFloor) *Monster {
	monsterType := rand.Intn(3) // 0 = Gobelin, 1 = Orc, 2 = Troll

	var m *Monster

	switch monsterType {
	case 0:
		m = &Monster{
			Nom:       "Gobelin",
			PVMax:     int(40 * floor.HPMult),
			PVActuels: int(40 * floor.HPMult),
			Attaque:   int(5 * floor.DmgMult),
			ExpReward: int(10 * floor.ExpMult),
			GoldMin:   int(3 * floor.GoldMult),
			GoldMax:   int(7 * floor.GoldMult),
		}
	case 1:
		m = &Monster{
			Nom:       "Orc",
			PVMax:     int(60 * floor.HPMult),
			PVActuels: int(60 * floor.HPMult),
			Attaque:   int(10 * floor.DmgMult),
			ExpReward: int(15 * floor.ExpMult),
			GoldMin:   int(5 * floor.GoldMult),
			GoldMax:   int(12 * floor.GoldMult),
		}
	case 2:
		m = &Monster{
			Nom:       "Troll",
			PVMax:     int(80 * floor.HPMult),
			PVActuels: int(80 * floor.HPMult),
			Attaque:   int(15 * floor.DmgMult),
			ExpReward: int(25 * floor.ExpMult),
			GoldMin:   int(10 * floor.GoldMult),
			GoldMax:   int(20 * floor.GoldMult),
		}
	}

	return m
}
