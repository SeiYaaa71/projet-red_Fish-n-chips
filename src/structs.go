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
	Gold             int
	Inventaire       []string
	Skill            []string
	Equipement       Equipment
	ExpActuelle      int
	ExpMax           int
	Niveau           int
	InventoryMax     int
	InventoryUpgrades int
	Effets []Effect
}

func (c *Character) isDead() bool {
	return c.PVActuels <= 0
}

type Monster struct {
	Nom       string
	PVMax     int
	PVActuels int
	Attaque   int
	ExpReward    int
	Effets []Effect
	GoldMin   int
	GoldMax   int
}

func (m *Monster) isDead() bool {
	return m.PVActuels <= 0
}
