package main

type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

type Character struct {
	Nom              string
	Classe           string
	PVMax            int
	PVActuels        int
	Gold             int
	Inventaire       []string
	Skill           []string
	Equipement       Equipment
	ExpActuelle      int
	ExpMax           int
	Niveau           int
	InventoryMax     int
	InventoryUpgrades int
}

type Monster struct {
	Nom       string
	PVMax     int
	PVActuels int
	Attaque   int
	ExpGagnee    int
}
