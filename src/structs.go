package main

type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

type Effect struct {
	Nom     string // Ex : "Poison"
	Tours   int    // Nombre de tours restants
	Degats  int    // Dégâts infligés par tour
	Cible   string // "player" ou "monster"
	Duration      int
	DamagePerTurn int
}


type Character struct {
	Nom              string
	Classe           string
	PVMax            int
	PVActuels        int
	Attaque		     int
	Gold             int
	Inventaire       []string
	Skill            []string
	Equipement       Equipment
	ExpActuelle      int
	ExpMax           int
	Potions		 	 int
	Niveau           int
	InventoryMax     int
	InventoryUpgrades int
	Effets 			 []Effect
	AbandonCombat    bool
}

func (c *Character) isDead() bool {
	return c.PVActuels <= 0
}

type Monster struct {
	Nom        string
	PVActuels  int
	PVMax      int
	Attaque     int
	GoldMin    int
	GoldMax    int
	ExpReward  int
	Effets     []Effect
}

func (m *Monster) isDead() bool {
	return m.PVActuels <= 0
}

type DungeonFloor struct {
	Numero   int     // Numéro de l’étage
	ExpMult  float64 // multiplicateur d’expérience
	GoldMult float64 // multiplicateur d’or
	HPMult   float64 // multiplicateur de PV du monstre
	DmgMult  float64 // multiplicateur de dégâts du monstre
}

func initOrc() *Monster {
	return &Monster{
		Nom:      "Orc",
		PVMax:    60,
		PVActuels: 60,
		Attaque:  10,
		ExpReward: 15,
		GoldMin:  5,
		GoldMax:  12,
	}
}

func initTroll() *Monster {
	return &Monster{
		Nom:      "Troll",
		PVMax:    80,
		PVActuels: 80,
		Attaque:  15,
		ExpReward: 25,
		GoldMin:  10,
		GoldMax:  20,
	}
}

func initGoblin() *Monster {
	return &Monster{
		Nom:       "Gobelin",
		PVMax:     40,
		PVActuels: 40,
		Attaque:   5,
		ExpReward: 10,
		GoldMin:   3,
		GoldMax:   7,
	}
}
