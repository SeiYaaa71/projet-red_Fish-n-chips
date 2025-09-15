package main

import "fmt"

// Applique les bonus de PV liés à l'équipement
func applyEquipmentBonus(c *Character) {
	// On réinitialise les PVMax de base en fonction de la classe
	switch c.Classe {
	case "Humain":
		c.PVMax = 100
	case "Elfe":
		c.PVMax = 80
	case "Nain":
		c.PVMax = 120
	}

	// Ajout des bonus selon l'équipement
	if c.Equipement.Tete == "Chapeau de l’aventurier" {
		c.PVMax += 10
	}
	if c.Equipement.Torse == "Tunique de l’aventurier" {
		c.PVMax += 25
	}
	if c.Equipement.Pieds == "Bottes de l’aventurier" {
		c.PVMax += 15
	}

	// Si PV actuels > PVMax après modif → ajustement
	if c.PVActuels > c.PVMax {
		c.PVActuels = c.PVMax
	}
}

// Fonction pour équiper un objet
func equipItem(c *Character, item string) {
	idx := -1
	for i, v := range c.Inventaire {
		if v == item {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println(Red + "❌ Vous n'avez pas cet objet dans votre inventaire." + Reset)
		return
	}

	// Retirer l'objet de l'inventaire
	c.Inventaire = append(c.Inventaire[:idx], c.Inventaire[idx+1:]...)

	rollback := func() { addToInventory(c, item) }

	switch item {
	case "Chapeau de l’aventurier":
		old := c.Equipement.Tete
		c.Equipement.Tete = item
		if old != "" && old != "Aucun" {
			if !addToInventory(c, old) {
				c.Equipement.Tete = old
				rollback()
				fmt.Println(Red + "❌ Inventaire plein ; équipement annulé." + Reset)
				return
			}
		}

	case "Tunique de l’aventurier":
		old := c.Equipement.Torse
		c.Equipement.Torse = item
		if old != "" && old != "Aucun" {
			if !addToInventory(c, old) {
				c.Equipement.Torse = old
				rollback()
				fmt.Println(Red + "❌ Inventaire plein ; équipement annulé." + Reset)
				return
			}
		}

	case "Bottes de l’aventurier":
		old := c.Equipement.Pieds
		c.Equipement.Pieds = item
		if old != "" && old != "Aucun" {
			if !addToInventory(c, old) {
				c.Equipement.Pieds = old
				rollback()
				fmt.Println(Red + "❌ Inventaire plein ; équipement annulé." + Reset)
				return
			}
		}

	default:
		rollback()
		fmt.Println(Red + "❌ Cet objet ne peut pas être équipé." + Reset)
		return
	}

	applyEquipmentBonus(c)
	fmt.Printf(Green+"✅ %s équipé avec succès !\n"+Reset, item)
	fmt.Printf("💖 PV : %d/%d\n", c.PVActuels, c.PVMax)
}
